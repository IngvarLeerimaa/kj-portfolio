package backend

import (
	"encoding/json"
	"time"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, client *Client) error

const (
	EventSendMessage    = "send_message"
	EventNewMessage     = "new_message"
	EventChangeRoom     = "change_room"
	EventGetCategories  = "get_categories"
	EventGetThreads     = "get_threads"
	EventPostThread     = "post_threads"
	EventGetComments    = "get_comments"
	EventAddComment     = "post_comment"
	EventGetOnlineUsers = "get_users"
	EventGetMessages    = "get_messages"
)

type AddCommentEvent struct {
	ThreadID int    `json:"threadID"`
	Title    string `json:"title"`
	Comment  string `json:"comment"`
}

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

type NewMessageEvent struct {
	SendMessageEvent
	Sent time.Time `json:"sent"`
}

type ChangeRoomEvent struct {
	UserID      int    `json:"userId"`
	LastMessage int    `json:"lastmessage"`
	Offset      int    `json:"offset"`
	RoomChange  bool   `json:"roomchange"`
	LastContent string `json:"lastContent"`
}

type GetMessagesEvent struct {
	Messages    []Message `json:"messages"`
	LastMessage int       `json:"lastmessage"`
}

type GetCategoriesEvent struct {
	Categories []byte `json:"categories"`
}

type GetThreadsEvent struct {
	Threads []byte `json:"threads"`
}
type GetOnlineUsersEvent struct {
	Users []OnlineUser `json:"users"`
}

type OnlineUser struct {
	Username    string    `json:"username"`
	UserId      int       `json:"userId"`
	LastMessage time.Time `json:"lastmessage"`
}
type UserNotInChat struct {
	UserId int `json:"userId"`
}
type PostThreadEvent struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category []string `json:"category"`
}

/* type RegisterEvent struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
*/
