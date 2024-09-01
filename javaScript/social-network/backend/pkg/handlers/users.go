package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/data/database"
	"social-network/pkg/helpers"
	"social-network/pkg/utils"
	"strconv"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (app *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	user := helpers.User{}
	user.Email = r.FormValue("email")
	exists, err := database.EmailExists(app.DB, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateUser, Error: %s\n", err)
		return
	}

	if exists != "" {
		response.Success = false
		response.Message = "Email already registered"
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("CreateUser, Error: %s\n", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
		return
	}

	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.DateOfBirth = r.FormValue("dateofbirth")
	user.Nickname = r.FormValue("nickname")
	user.About = r.FormValue("about")

	password, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateUser, Error: %s\n", err)
		return
	}
	user.Password = string(password)

	r.ParseMultipartForm(10 << 20)
	user.Avatar = utils.SaveImage(r.FormFile("avatar"))

	if user.Avatar == "" {
		user.Avatar = "avatar.png"
	}
	err = database.SaveUserDetails(app.DB, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateUser, Error: %s\n", err)
		return
	}

	response.Success = true
	response.Message = "Account registered, please login"
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateUser, Error: %s\n", err)
		return
	}
	w.Write(jsonResponse)
}

func (app *App) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetUser, Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Printf("GetUser, Error: %s\n", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetUser, Error: %s\n", err)
		return
	}

	user, err := database.User(app.DB, userID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetUser, Error: %s\n", err)
		return
	}

	user.Following, user.Pending, user.Follower, err = database.FollowData(app.DB, userID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetUser, Error: %s\n", err)
		return
	}

	if !user.CurrentUser && !user.Public && !user.Following {
		user.Email, user.DateOfBirth, user.About, user.Nickname = "", "", "", ""
	}

	response := struct {
		User helpers.User `json:"user"`
	}{
		User: user,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetUser, Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}

func (app *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Println("GetUsers, Error getting session ID:", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetUsers, Error: %s\n", err)
		return
	}

	users, err := database.UserList(app.DB, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetUsers, Error: %s\n", err)
		return
	}

	for i := 0; i < len(users); i++ {
		users[i].Following, users[i].Pending, users[i].Follower, err = database.FollowData(app.DB, users[i].ID, currentUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("GetUsers, Error: %s\n", err)
			return
		}
	}

	response := struct {
		Users []helpers.User `json:"users"`
	}{
		Users: users,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetUsers, Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}

func (app *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	privacy := struct {
		Privacy string `json:"privacy"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&privacy); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("UpdateUser, Error: %s\n", err)
		return
	}

	privacyVal, err := strconv.Atoi(privacy.Privacy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("UpdateUser, Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("UpdateUser, Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("UpdateUser, Error: %s\n", err)
		return
	}

	err = database.UpdateUser(app.DB, userID, privacyVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("UpdateUser, Error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var credentials helpers.User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Login, Error: %s\n", err)
		return
	}

	passHash, err := database.EmailExists(app.DB, credentials.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Login, Error: %s\n", err)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(credentials.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid credentials")
		return
	}
	sessionID := uuid.Must(uuid.NewV4()).String()
	userID, err := database.SaveSession(app.DB, sessionID, credentials.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Login, Error: %s\n", err)
		return
	}

	response := helpers.SessionResponse{
		SessionID: sessionID,
		UserID:    userID,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Login, Error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (app *App) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Println("Error getting session ID:", err)
		return
	}

	database.DeleteSession(app.DB, sCookie.Value)
}

func (app *App) CheckSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := struct {
		UserID int `json:"userid"`
	}{
		UserID: 0,
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("CheckSession, Error: %s\n", err)
			return
		}
		w.Write(jsonResponse)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CheckSession, Error: %s\n", err)
		return
	}

	if userID < 1 {
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("CheckSession, Error: %s\n", err)
			return
		}
		w.Write(jsonResponse)
		return
	}

	response.UserID = userID

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CheckSession, Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}

func (app *App) GetNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Printf("GetNotifications, Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetNotifications, Error: %s\n", err)
		return
	}

	response, err := database.Notifications(app.DB, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetNotifications, Error: %s\n", err)
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetNotifications, Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}
