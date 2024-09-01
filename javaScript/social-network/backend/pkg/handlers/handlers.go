package handlers

import (
	"database/sql"
	"net/http"
)

type App struct {
	DB *sql.DB
	clientManager *Manager
}

func NewApp(db *sql.DB) *App {
	return &App{
		DB: db,
		clientManager: NewManager(),
	}
}

func AddHandlers(router *http.ServeMux, db *sql.DB) {
	app := NewApp(db)

	router.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./data/images/"))))

	router.HandleFunc("/ws", app.WebSocketHandler)

	router.HandleFunc("/api/v1/user/register", app.CreateUser)
	router.HandleFunc("/api/v1/user/login", app.Login)
	router.HandleFunc("/api/v1/user/logout", app.Logout)
	router.HandleFunc("/api/v1/user/session", app.CheckSession)
	router.HandleFunc("/api/v1/user/update", app.UpdateUser)
	router.HandleFunc("/api/v1/user/follow", app.Follow)
	router.HandleFunc("/api/v1/follow/update", app.UpdateFollow)

	router.HandleFunc("/api/v1/user", app.GetUser)
	router.HandleFunc("/api/v1/user/followers", app.getUserFollow)
	router.HandleFunc("/api/v1/user/notifications", app.GetNotifications)
	router.HandleFunc("/api/v1/users", app.GetUsers)

	router.HandleFunc("/api/v1/posts", app.GetPosts)
	router.HandleFunc("/api/v1/posts/user", app.GetUserPosts)
	router.HandleFunc("/api/v1/post/create", app.CreatePost)

	router.HandleFunc("/api/v1/comment/create", app.CreateComment)
	router.HandleFunc("/api/v1/comments", app.GetComments)

	router.HandleFunc("/api/v1/group/create", app.CreateGroup)
	router.HandleFunc("/api/v1/group", app.GetGroup)
	router.HandleFunc("/api/v1/groups", app.GetGroups)
	router.HandleFunc("/api/v1/group/user/create", app.CreateGroupUser)
	router.HandleFunc("/api/v1/group/user/update", app.UpdateGroupUser)

	router.HandleFunc("/api/v1/message/create", app.CreateMessage)
	router.HandleFunc("/api/v1/message", app.GetLastMessage)
	router.HandleFunc("/api/v1/messages", app.GetMessages)
	router.HandleFunc("/api/v1/group/message", app.GetLastGroupMessage)
	router.HandleFunc("/api/v1/group/messages", app.GetGroupMessages)

	router.HandleFunc("/api/v1/event/create", app.CreateEvent)
	router.HandleFunc("/api/v1/event/going", app.EventUser)
}
