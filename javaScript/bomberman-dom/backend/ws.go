package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	mutex      sync.Mutex
	name       string
	room       *Room
	num        int
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

type Message struct {
	MessageType string          `json:"messageType"`
	Message     string          `json:"message"`
	X           int             `json:"x"`
	Y           int             `json:"y"`
	Cells       []ExplosionCell `json:"cells"`
}

func (m *Manager) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &Client{
		connection: conn,
		manager:    m,
		name:       name,
	}

	m.addClient(client)
	if len(AvailableRooms) == 0 {
		room := CreateRoom()
		Rooms[room.ID] = room
		AvailableRooms[room.ID] = room
		client.room = room
	} else {
		for _, r := range AvailableRooms {
			client.room = r
			break
		}
	}
	// go routine to listen and read messages per client
	go client.ReadMessages()
	go client.WriteMessages()

	client.room.AddUser(client)
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
		client.room.RemoveUser(client)
	}
}

func (c *Client) ReadMessages() {
	defer func() {
		// clean up connection
		c.manager.removeClient(c)
	}()
	c.connection.SetReadLimit(5096)
	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}
	c.connection.SetPongHandler(c.pongHandler)
	for {
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}

		/* fmt.Println("we are here", messageType, string(payload)) */

		var msg Message
		if err := json.Unmarshal(payload, &msg); err != nil {
			log.Println("error unmarshalling message", err)
			continue // Continue the loop instead of returning.
		}
		if msg.MessageType == "message" {
			c.room.BroadcastMessage(c, msg.Message)
		}
		if msg.MessageType == "location" {
			c.room.SendPlayerLocation(c, msg.X, msg.Y, msg.Message)
		}
		if msg.MessageType == "bomb" {
			c.room.SendBomb(msg.X, msg.Y)
		}
		if msg.MessageType == "explosionCells" {
			c.room.SendExplosion(msg.Cells)
		}
		if msg.MessageType == "powerup" {
			c.room.SendRemovePowerUp(msg.X)
		}
		if msg.MessageType == "dead" {
			c.room.SendRemovePlayer(c.num)
			c.room.removeAlive(c.num)
		if len(c.room.Alive) == 1 {
			c.room.SendWinner()
		}
		}
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
