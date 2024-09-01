package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"social-network/data/database"
	"social-network/pkg/helpers"
	"strconv"
	"time"
)

func (app *App) CreateMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	message := helpers.Message{}

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error: %s\n", err)
		return
	}

	userMessage := true
	if message.FromID == 0 {
		userMessage = false
	}

	message.FromID = userID
	message.Created = time.Now().Format("2006-01-02 15:04:05")
	if userMessage {
		if err := database.AddMessage(app.DB, message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error: %s\n", err)
			return
		}
	
		if err := app.SendChatMessage(message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error: %s\n", err)
			return
		}
	} else {
		if err := database.AddGroupMessage(app.DB, message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error: %s\n", err)
			return
		}
	
		if err := app.SendGroupChatMessage(message, userID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error: %s\n", err)
			return
		}
	}
	

	w.WriteHeader(http.StatusOK)
}

func (app *App) GetLastMessage(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Error: %s\n", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	message, err := database.LastMessage(app.DB, userID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	response := struct {
		Message helpers.Message `json:"message"`
	}{
		Message: message,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}

func (app *App) GetLastGroupMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	groupID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetLastGroupMessage, Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Printf("GetLastGroupMessage, Error: %s\n", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetLastGroupMessage, Error: %s\n", err)
		return
	}

	isMember, err := database.IsGroupMember(app.DB, groupID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetLastGroupMessage, Error: %s\n", err)
		return
	}
	if !isMember {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}

	message, err := database.LastGroupMessage(app.DB, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetLastGroupMessage, Error: %s\n", err)
		return
	}

	response := struct {
		Message helpers.Message `json:"message"`
	}{
		Message: message,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetLastGroupMessage, Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}

func (app *App) GetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Println("Error getting session ID:", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	messages, err := database.Messages(app.DB, userID, currentUser, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	response := struct {
		Messages []helpers.Message `json:"messages"`
	}{
		Messages: messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}

func (app *App) GetGroupMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Println("Error getting session ID:", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}


	groupID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	isMember, err := database.IsGroupMember(app.DB, groupID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetLastGroupMessage, Error: %s\n", err)
		return
	}
	if !isMember {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	messages, err := database.GroupMessages(app.DB, groupID, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	response := struct {
		Messages []helpers.Message `json:"messages"`
	}{
		Messages: messages,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}
