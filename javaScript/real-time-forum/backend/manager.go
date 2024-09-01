package backend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	portNum string
)

type Manager struct {
	clients ClientList
	sync.RWMutex

	otps RetentionMap

	handlers map[string]EventHandler
}

func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
		otps:     NewRetentionMap(ctx, 5*time.Second),
	}
	m.setupEventHandlers()
	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessageHandler
	m.handlers[EventChangeRoom] = ChatRoomHandler
	m.handlers[EventGetCategories] = GetCategoriesHandler
	m.handlers[EventGetThreads] = GetThreadsHandler
	m.handlers[EventGetOnlineUsers] = GetOnlineUsersHandler
	m.handlers[EventGetComments] = GetCommentsHandler
	m.handlers[EventPostThread] = PostThreadHandler
	m.handlers[EventAddComment] = AddCommentHandler
}

func AddCommentHandler(event Event, c *Client) error {
	// Marshal Payload into wanted format
	var comment []string
	//fmt.Println("event.Payload: ", string(event.Payload))
	if err := json.Unmarshal(event.Payload, &comment); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}
	//fmt.Println("commentevent: ", comment)
	var commentevent AddCommentEvent
	tra, err := strconv.Atoi(comment[0])
	if err != nil {
		return fmt.Errorf("failed to convert ThreadId to int: %v", err)

	}
	commentevent.ThreadID = tra
	commentevent.Title = string(comment[1])
	commentevent.Comment = string(comment[2])

	//fmt.Println("commentevent: ", commentevent)
	// Prepare an Outgoing Message to others
	ok := 1
	if commentevent.Title == "" || commentevent.Comment == "" {
		ok = 0
		return fmt.Errorf("failed to create comment")
	}
	AddComment(commentevent.ThreadID, c.username, commentevent.Title, commentevent.Comment)
	data, err := json.Marshal(ok)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	// Place payload into an Event
	outgoingEvent := Event{
		Payload: data,
		Type:    EventAddComment,
	}
	//fmt.Println("outgoingEvent: ", outgoingEvent)
	// Broadcast to all other Clients
	c.egress <- outgoingEvent
	return nil
}

func GetCommentsHandler(event Event, c *Client) error {
	// Marshal Payload into wanted format
	var idString string
	//fmt.Println("event.Payload: ", string(event.Payload))
	if err := json.Unmarshal(event.Payload, &idString); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return fmt.Errorf("failed to convert ThreadId to int: %v", err)
	}
	//fmt.Println("id mis läheb Addcomments: ", id)
	data.AddComments(id)
	//fmt.Println("data.Comments: ", data.Comments)
	// Marshal posts into JSON format
	commentsJSON, err := json.Marshal(data.Comments)
	if err != nil {
		return fmt.Errorf("failed to marshal posts: %v", err)
	}

	// Place payload into an Event
	outgoingEvent := Event{
		Type:    EventGetComments,
		Payload: commentsJSON,
	}

	// Send to the client
	c.egress <- outgoingEvent
	return nil
}

func PostThreadHandler(event Event, c *Client) error {
	// Marshal Payload into wanted format
	var postevent PostThreadEvent
	if err := json.Unmarshal(event.Payload, &postevent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}
	//fmt.Println("postevent: ", postevent)

	res := CreateThread(c.username, postevent)
	// Prepare an Outgoing Message to others

	data, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	// Place payload into an Event
	outgoingEvent := Event{
		Payload: data,
		Type:    EventPostThread,
	}

	// Broadcast to all other Clients
	c.egress <- outgoingEvent
	return nil
}

func GetOnlineUsersHandler(event Event, c *Client) error {
	// Prepare a slice to store the usernames

	var users []OnlineUser

	clients := c.manager.clients
	// Iterate over the client list and extract usernames
	for client := range clients {
		for clnt := range clients {
			if client.userID != clnt.userID {
				lamp := OnlineUser{}
				lamp.Username = clnt.username
				lamp.UserId = clnt.userID
				if date, ok := GetLastMessageDate(client.userID, clnt.userID); ok {
					lamp.LastMessage = date
				}
				/* fmt.Println("user", lamp.Username)
				fmt.Println("date:", lamp.LastMessage.IsZero()) */
				users = append(users, lamp)
			}
		}
		//siia on vaja loogikat, et ta saadaks tagasi, et mitte kedagi pole online.
		//Siis ei jää clientile ka ette viimati saadetud list
		// või muuta lõppus kui users on tühi, siis saadab tühja listi
		if len(users) == 0 {
			return nil
		}
		sort.SliceStable(users, func(i, j int) bool {
			if !users[i].LastMessage.IsZero() && !users[j].LastMessage.IsZero() {
				return users[i].LastMessage.After(users[j].LastMessage)
			}
			if !users[i].LastMessage.IsZero() {
				return true
			}
			if !users[j].LastMessage.IsZero() {
				return false
			}
			return users[i].Username < users[j].Username
		})
		// Create an instance of OnlineUsers struct
		onlineUsers := GetOnlineUsersEvent{
			Users: users,
		}

		// Marshal onlineUsers into JSON format
		usersJSON, err := json.Marshal(onlineUsers)
		if err != nil {
			return fmt.Errorf("failed to marshal online users: %v", err)
		}

		// Place payload into an Event
		outgoingEvent := Event{
			Type:    EventGetOnlineUsers,
			Payload: usersJSON,
		}

		// Send to the client
		client.egress <- outgoingEvent
		users = []OnlineUser{}
	}

	return nil
}

func ChatRoomHandler(event Event, c *Client) error {

	var changeRoomEvent ChangeRoomEvent
	//fmt.Println("event.Payload: ", string(event.Payload))
	if err := json.Unmarshal(event.Payload, &changeRoomEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}
	//fmt.Println(changeRoomEvent)

	if !changeRoomEvent.RoomChange {
		c.chatroom = -1
		return nil
	}

	// Add Client to chat room
	offset := changeRoomEvent.Offset
	if offset == 0 {
		c.chatroom = changeRoomEvent.UserID
	}

	messages := GetMessages(c.userID, c.chatroom, offset)
	var lastMesRep = false
	var lastMesRepIndex int
	for i := 1; i < len(messages)-1; i++ {
		if messages[i] == messages[i-1] {
			lastMesRep = true
			lastMesRepIndex = i
			break
		}
	}
	if lastMesRep {
		messages[lastMesRepIndex].From = ""
		messages[lastMesRepIndex].Content = "You've reached the end of the chat"
		messages[lastMesRepIndex].Created = time.Now()
		messages = messages[:lastMesRepIndex+1]
	}
	offset += 10

	data := GetMessagesEvent{
		Messages:    messages,
		LastMessage: offset,
	}

	outgoingEventJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal past messages: %v", err)
	}

	outgoingEvent := Event{
		Type:    EventGetMessages,
		Payload: outgoingEventJSON,
	}

	c.egress <- outgoingEvent

	return nil
}

func GetThreadsHandler(event Event, c *Client) error {
	if string(event.Payload) == "" {

		asd := ""
		threadsJSON, err := json.Marshal(asd)
		if err != nil {
			return fmt.Errorf("failed to marshal posts: %v", err)
		}
		outgoingEvent := Event{
			Type:    "",
			Payload: threadsJSON,
		}

		// Send to the client
		c.egress <- outgoingEvent
		return nil
	}
	//Probably better way to do this
	var idString string
	if err := json.Unmarshal(event.Payload, &idString); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	//fmt.Println("threads in getThreadsHandler: ", idString)

	id, err := strconv.Atoi(idString)
	if err != nil {
		return fmt.Errorf("failed to convert ThreadId to int: %v", err)
	}

	data.AddThreads(id)
	// Marshal posts into JSON format
	threadsJSON, err := json.Marshal(data.Threads)
	if err != nil {
		return fmt.Errorf("failed to marshal posts: %v", err)
	}
	//fmt.Println("data: ", string(threadsJSON))
	// Place payload into an Event
	outgoingEvent := Event{
		Type:    EventGetThreads,
		Payload: threadsJSON,
	}

	// Send to the client
	c.egress <- outgoingEvent
	return nil

}

func GetCategoriesHandler(event Event, c *Client) error {
	data.AddCategories()
	// Prepare categories as a string
	//categories := "suvaline, suvaline2, suvaline3"

	// Marshal categories into JSON format
	categoriesJSON, err := json.Marshal(data.Categories)
	if err != nil {
		return fmt.Errorf("failed to marshal categories: %v", err)
	}

	//fmt.Println("data: ", string(categoriesJSON))

	// Place payload into an Event
	outgoingEvent := Event{
		Type:    EventGetCategories,
		Payload: categoriesJSON,
	}

	// Send to the client
	c.egress <- outgoingEvent
	return nil
}

func SendMessageHandler(event Event, c *Client) error {
	// Marshal Payload into wanted format
	var chatevent SendMessageEvent
	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	/* fmt.Println("chatevent: ", chatevent)
	fmt.Println("Client on:", c) */
	AddMessage(c.userID, c.chatroom, chatevent.Message)
	// Prepare an Outgoing Message to others
	var broadMessage NewMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatevent.Message
	broadMessage.From = c.username

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	// Place payload into an Event
	outgoingEvent := Event{
		Payload: data,
		Type:    EventNewMessage,
	}

	/* var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewMessage */
	// Broadcast to all other Clients
	for client := range c.manager.clients {
		// Only send to clients inside the same chatroom
		if client.userID == c.chatroom || client.userID == c.userID {
			if client.userID != c.userID && client.chatroom != c.userID {
				idData, err := json.Marshal(UserNotInChat{c.userID})
				if err != nil {
					return fmt.Errorf("failed to marshal broadcast message: %v", err)
				}
				notInChat := Event{
					Payload: idData,
					Type:    EventNewMessage,
				}
				client.egress <- notInChat
			} else {
				client.egress <- outgoingEvent
			}

		}

	}
	return nil
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("there is no such event type")
	}
}

func (m *Manager) serveWs(w http.ResponseWriter, r *http.Request) {

	otp := r.URL.Query().Get("otp")
	username := r.URL.Query().Get("username")
	if otp == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !m.otps.VerifyOTP(otp) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("WebSocket connected")
	// Upgrade the HTTP connection to a websocket connection
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)
	client.username = username
	client.userID = uid
	m.addClient(client)

	// go routine to listen and read messages per client
	go client.ReadMessages()
	go client.WriteMessages()
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true

}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, alreadyConnected := m.clients[client]; alreadyConnected {
		client.connection.Close()
		delete(m.clients, client)
	}
}
func setPortNum(p string) {
	portNum = p
}

func checkOrigin(r *http.Request) bool {

	// Grab the request origin
	origin := r.Header.Get("Origin")

	switch origin {
	case "https://localhost:" + portNum:
		return true
	default:
		return false
	}
}

func (m *Manager) sessionHandler(w http.ResponseWriter, r *http.Request) {
	type sessionRequest struct {
		SessionID string `json:"sessionID"`
	}
	var req sessionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println(CheckSession(req.SessionID))
	if _, username, ok := CheckSession(req.SessionID); ok {
		type response struct {
			OTP      string `json:"otp"`
			Username string `json:"username"`
			Redirect string `json:"redirect"`
			UserID   int    `json:"userID"`
		}

		otp := m.otps.NewOTP()
		resp := response{
			OTP:      otp.Key,
			Username: username,
			UserID:   uid,
			Redirect: "/",
		}
		data, err := json.Marshal(resp)
		if err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	} else {
		return
	}
}
