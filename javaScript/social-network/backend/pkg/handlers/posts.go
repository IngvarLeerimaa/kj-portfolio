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
	"strings"
)

func (app *App) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	var post helpers.Post
	groupID := 0

	post.UserID = userID
	post.Text = r.FormValue("text")

	post.Privacy, err = strconv.Atoi(r.FormValue("privacy"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	if post.Privacy == 2 {
		followers := strings.Split(r.FormValue("followers"), ",")

		for _, f := range followers {
			follower, err := strconv.Atoi(f)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("Error: %s\n", err)
				return
			}
			post.Followers = append(post.Followers, follower)
		}
	}

	if post.Privacy == 3 {
		groupID, err = strconv.Atoi(r.FormValue("groupid"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error: %s\n", err)
			return
		}
		isGroupMember, err := database.IsGroupMember(app.DB, groupID, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error: %s\n", err)
			return
		}

		if !isGroupMember {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
	}

	r.ParseMultipartForm(10 << 20)
	post.Image = utils.SaveImage(r.FormFile("image"))

	if err = database.SavePost(app.DB, post, groupID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	var comment helpers.Comment

	comment.UserID = userID
	comment.Text = r.FormValue("text")

	comment.PostID, err = strconv.Atoi(r.FormValue("postId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	r.ParseMultipartForm(10 << 20)
	comment.Image = utils.SaveImage(r.FormFile("image"))

	if err = database.SaveComment(app.DB, comment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *App) GetComments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}
	postID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	comments, err := database.Comments(app.DB, userID, postID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	response := struct {
		Comments []helpers.Comment `json:"comments"`
	}{
		Comments: comments,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}
func (app *App) GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sCookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	userID, err := database.ValidateSessionID(app.DB, sCookie.Value)
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

	posts, err := database.Posts(app.DB, userID, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	response := struct {
		Posts []helpers.Post `json:"posts"`
	}{
		Posts: posts,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}

func (app *App) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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

	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
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
		http.Error(w, "not authorized to get this users posts", http.StatusUnauthorized)
		return
	}

	posts, err := database.UserPosts(app.DB, userID, currentUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	response := struct {
		Posts []helpers.Post `json:"posts"`
	}{
		Posts: posts,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %s\n", err)
		return
	}

	w.Write(jsonResponse)
}
