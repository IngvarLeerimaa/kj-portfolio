package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/data/database"
	"social-network/pkg/helpers"
	"strconv"
	"strings"
)

func (app *App) CreateGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateGroup, Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateGroup, Error: %s\n", err)
		return
	}

	var group helpers.Group

	title := r.FormValue("title")

	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	title = strings.ToUpper(title)
	title = strings.TrimSpace(title)

	desc := r.FormValue("description")
	desc = strings.TrimSpace(desc)

	group.AdminID = userID
	group.Title = title
	group.Description = desc

	if err = database.SaveGroup(app.DB, group); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateGroup, Error: %s\n", err)
		return
	}
	fmt.Printf("Group created with Title: %v, Description: %v and AdminID: %d\n", group.Title, group.Description, group.AdminID)

}

func (app *App) GetGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	groupID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroup, Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Printf("GetGroup, Error: %s\n", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroup, Error: %s\n", err)
		return
	}

	isGroupMember, err := database.IsGroupMember(app.DB, groupID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroup, Error: %s\n", err)
		return
	}

	if !isGroupMember {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	posts, err := database.GroupPosts(app.DB, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroup, Error: %s\n", err)
		return
	}

	events, err := database.GroupEvents(app.DB, groupID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroup, Error: %s\n", err)
		return
	}

	members, invited, err := database.MembersData(app.DB, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroup, Error: %s\n", err)
		return
	}

	response := struct {
		Posts   []helpers.Post       `json:"posts"`
		Events  []helpers.GroupEvent `json:"events"`
		Members []int                `json:"members"`
		Invited []int                `json:"invited"`
	}{
		Posts:   posts,
		Events:  events,
		Members: members,
		Invited: invited,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroup, Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}

func (app *App) GetGroups(w http.ResponseWriter, r *http.Request) {
	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Println("GetGroups, Error getting session ID:", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroups, Error: %s\n", err)
		return
	}

	groups, err := database.Groups(app.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroups, Error: %s\n", err)
		return
	}

	for i := 0; i < len(groups); i++ {
		groups[i].Joined, groups[i].Invited, groups[i].Requested, err = database.JoinData(app.DB, groups[i].GroupID, currentUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("GetGroups, Error: %s\n", err)
			return
		}
	}

	response := struct {
		Groups []helpers.Group `json:"groups"`
	}{
		Groups: groups,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("GetGroups, Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}

func (app *App) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	event := helpers.GroupEvent{}

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("CreateEvent, Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		log.Printf("CreateEvent, Error: %s\n", err)
		return
	}

	currentUser, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateEvent, Error: %s\n", err)
		return
	}

	isGroupMember, err := database.IsGroupMember(app.DB, event.GroupID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateEvent, Error: %s\n", err)
		return
	}

	if !isGroupMember {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	event.CreatorID = currentUser

	eventID, err := database.AddEvent(app.DB, event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateEvent, Error: %s\n", err)
		return
	}

	event.EventID = int(eventID)

	if err = database.AddEventUser(app.DB, event.EventID, currentUser, event.Going); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateEvent, Error: %s\n", err)
		return
	}

	n := helpers.Notification{
		NotificationType: "event",
		GroupEvent:       event,
	}

	rec, err := database.GroupUsers(app.DB, event.GroupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateEvent, Error: %s\n", err)
		return
	}

	recipients := []int{}

	for _, v := range rec {
		if v != currentUser {
			recipients = append(recipients, v)
		}
	}

	err = app.Notification(n, recipients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateEvent, Error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) EventUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("EventUser, Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("EventUser, Error: %s\n", err)
		return
	}

	data := struct {
		Going   bool `json:"going"`
		EventID int  `json:"eventId"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("EventUser, Error: %s\n", err)
		return
	}

	groupID, err := database.GetGroupID(app.DB, data.EventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("EventUser, Error: %s\n", err)
		return
	}

	isGroupMember, err := database.IsGroupMember(app.DB, groupID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("EventUser, Error: %s\n", err)
		return
	}

	if !isGroupMember {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		log.Printf("EventUser, Error: %s\n", err)
		return
	}

	err = database.AddEventUser(app.DB, data.EventID, userID, data.Going)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("EventUser, Error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) CreateGroupUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	data := struct {
		GroupID int  `json:"groupId"`
		UserID  int  `json:"userId"`
		Invite  bool `json:"invite"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("CreateGroupUser, Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateGroupUser, Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("CreateGroupUser, Error: %s\n", err)
		return
	}

	if data.Invite {
		isMember, err := database.IsGroupMember(app.DB, data.GroupID, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("CreateGroupUser, Error: %s\n", err)
			return
		}
		if !isMember {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		err = database.AddGroupUser(app.DB, data.GroupID, data.UserID, 1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("CreateGroupUser, Error: %s\n", err)
			return
		}
		n := helpers.Notification{
			NotificationType: "invite",
			GroupID:          data.GroupID,
		}
		err = app.Notification(n, []int{data.UserID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("CreateGroupUser, Error: %s\n", err)
			return
		}
	} else {
		err = database.AddGroupUser(app.DB, data.GroupID, userID, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("CreateGroupUser, Error: %s\n", err)
			return
		}

		adminID, err := database.GroupAdminID(app.DB, data.GroupID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("CreateGroupUser, Error: %s\n", err)
			return
		}
		n := helpers.Notification{
			NotificationType: "request",
			UserID:           userID,
			GroupID:          data.GroupID,
		}
		err = app.Notification(n, []int{adminID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("CreateGroupUser, Error: %s\n", err)
			return
		}

	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) UpdateGroupUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := struct {
		GroupID int  `json:"groupId"`
		UserID  int  `json:"userId"`
		Request bool `json:"request"`
		Confirm bool `json:"confirm"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("UpdateGroupUser, Error: %s\n", err)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("UpdateGroupUser, Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("UpdateGroupUser, Error: %s\n", err)
		return
	}

	if data.Request {
		adminID, err := database.GroupAdminID(app.DB, data.GroupID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("UpdateGroupUser, Error: %s\n", err)
			return
		}
		if adminID != userID {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			log.Printf("UpdateGroupUser, Error: %s\n", err)
			return
		}
		if data.Confirm {
			err = database.ConfirmGroupUser(app.DB, data.GroupID, data.UserID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("UpdateGroupUser, Error: %s\n", err)
				return
			}
		} else {
			err = database.DeleteGroupUser(app.DB, data.GroupID, data.UserID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("UpdateGroupUser, Error: %s\n", err)
				return
			}
		}
	} else {
		if data.Confirm {
			err = database.ConfirmGroupUser(app.DB, data.GroupID, userID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("UpdateGroupUser, Error: %s\n", err)
				return
			}
		} else {
			err = database.DeleteGroupUser(app.DB, data.GroupID, userID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("UpdateGroupUser, Error: %s\n", err)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}
