package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/data/database"
	"social-network/pkg/helpers"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	userID     int
}
type ClientList map[*Client]bool

type Manager struct {
	clients ClientList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

func (a *App) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	m := a.clientManager
	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Upgrade the HTTP connection to a websocket connection
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &Client{
		connection: conn,
		manager:    m,
		userID:     userID,
	}

	m.addClient(client)

	// go routine to listen and read messages per client
	go client.ReadMessages()
	go client.WriteMessages()
	a.UserOnline(userID)
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
	defer UserOffline(m.clients, client.userID)
}

func (c *Client) ReadMessages() {
	defer func() {
		//clean up connection
		c.manager.removeClient(c)
	}()
	c.connection.SetReadLimit(512)
	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}
	c.connection.SetPongHandler(c.pongHandler)
	for {
		messageType, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}
		fmt.Println(messageType, string(payload))
	}
}
func (c *Client) WriteMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()
	ticker := time.NewTicker(pingInterval)
	defer ticker.Stop()

	for range ticker.C {
		if err := c.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			return
		}
	}
}
func (c *Client) pongHandler(pongMsg string) error {
	//checkib kas on elus
	/* log.Println("pong") */
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

func (a *App) SendChatMessage(msg helpers.Message) error {
	for client := range a.clientManager.clients {
		if client.userID == msg.ToID {
			data := struct {
				Type    string          `json:"type"`
				Message helpers.Message `json:"message"`
			}{
				Type:    "newMessage",
				Message: msg,
			}
			msgJSON, err := json.Marshal(data)
			if err != nil {
				return err
			}

			err = client.connection.WriteMessage(1, msgJSON)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *App) SendGroupChatMessage(msg helpers.Message, currentUser int) error {
	groupUsers, err := database.GroupUsers(a.DB, msg.ToID)
	if err != nil {
		return err
	}
	data := struct {
		Type    string          `json:"type"`
		Message helpers.Message `json:"message"`
	}{
		Type:    "newGroupMessage",
		Message: msg,
	}
	msgJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for client := range a.clientManager.clients {
		for _, id := range groupUsers {
			if client.userID == id && client.userID != currentUser {
				err = client.connection.WriteMessage(1, msgJSON)
				if err != nil {
					return err
				}
			}
		}
		
	}
	return nil
}

func (a *App) UserOnline(currentUser int) error {
	data := struct {
		Type   string `json:"type"`
		UserID int    `json:"userId"`
		Online bool   `json:"online"`
	}{
		Type:   "online",
		Online: true,
	}
	for client := range a.clientManager.clients {
		if client.userID == currentUser {
			for c := range a.clientManager.clients {
				if c.userID != currentUser {
					data.UserID = c.userID
					onlineJSON, err := json.Marshal(data)
					if err != nil {
						return err
					}

					err = client.connection.WriteMessage(1, onlineJSON)
					if err != nil {
						return err
					}
				}
			}
		} else {
			data.UserID = currentUser
			onlineJSON, err := json.Marshal(data)
			if err != nil {
				return err
			}

			err = client.connection.WriteMessage(1, onlineJSON)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UserOffline(clients ClientList, currentUser int) error {
	data := struct {
		Type   string `json:"type"`
		UserID int    `json:"userId"`
		Online bool   `json:"online"`
	}{
		Type:   "online",
		UserID: currentUser,
		Online: false,
	}
	offlineJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for client := range clients {
		err = client.connection.WriteMessage(1, offlineJSON)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) Notification(n helpers.Notification, recipients []int) error {
	data := struct {
		Type     string `json:"type"`
		Notification helpers.Notification `json:"notification"`
	}{
		Type:     "notification",
		Notification: n,
	}

	dJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for client := range a.clientManager.clients {
		for _, id := range recipients {
			if client.userID == id {
				err = client.connection.WriteMessage(1, dJSON)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *App) SendFollowDecision(userID, currentUser int, decision, request bool) error {
	for client := range a.clientManager.clients {
		if client.userID == userID {

			data := struct {
				Type     string `json:"type"`
				UserID   int    `json:"userId"`
				Decision bool   `json:"decision"`
				Request  bool   `json:"request"`
			}{
				Type:     "follow",
				UserID:   currentUser,
				Decision: decision,
				Request:  request,
			}
			dJSON, err := json.Marshal(data)
			if err != nil {
				return err
			}

			err = client.connection.WriteMessage(1, dJSON)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
