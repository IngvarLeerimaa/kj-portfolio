package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"social-network/data/database"
	"social-network/pkg/helpers"
	"strconv"
)

func (app *App) Follow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	follow := struct {
		UserID int  `json:"userId"`
		Follow bool `json:"follow"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	if follow.Follow {
		err = database.AddFollow(app.DB, follow.UserID, currentUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error: %s\n", err)
			return
		}
	} else {
		err = database.DeleteFollow(app.DB, follow.UserID, currentUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error: %s\n", err)
			return
		}
	}

	err = app.SendFollowDecision(follow.UserID, currentUser, follow.Follow, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) UpdateFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	follow := struct {
		UserID int  `json:"userId"`
		Decision bool `json:"decision"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&follow); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	if follow.Decision{
		err = database.ConfirmFollow(app.DB, follow.UserID, currentUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error: %s\n", err)
			return
		}
	} else {
		err = database.DeleteFollow(app.DB, currentUser, follow.UserID)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}
	}
	err = app.SendFollowDecision(follow.UserID, currentUser, follow.Decision, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) getUserFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	user, err := database.User(app.DB, userID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	following, _, _, err := database.FollowData(app.DB, userID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	if !user.Public && !user.CurrentUser && !following {
		http.Error(w, "not authorized to get data", http.StatusUnauthorized)
		return
	}

	flwrs, err := database.Followers(app.DB, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	flwng, err := database.Following(app.DB, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	response := struct {
		Followers []helpers.User `json:"followers"`
		Following []helpers.User `json:"following"`
	}{
		Followers: flwrs,
		Following: flwng,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}
