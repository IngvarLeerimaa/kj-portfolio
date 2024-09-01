package backend

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait = 10 * time.Second

	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan Event
	chatroom   int
	username   string
	userID     int
}

type ClientList map[*Client]bool

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan Event),
	}
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
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}

		var request Event

		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			break
		}
		//fmt.Println("request: ", request)
		if err := c.manager.routeEvent(request, c); err != nil {
			log.Printf("Error routing message: %v", err)
			break
		}

	}
}

func (c *Client) WriteMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	ticker := time.NewTicker(pingInterval)

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Printf("connection closed: %v", err)
				}
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Printf("failed to marshal message: %v", err)
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Printf("failed to send message: %v", err)
				return
			}
			//log.Printf("message sent")

		case <-ticker.C:
			//checkib kas on elus
			/* log.Println("ping") */
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				/* log.Println("writemsg: ", err)
				log.Println("client ln 107 tickeriga vbl on vaja case kus tuleb ws restartida..")
				log.Println("Kui urlil enter panna siis ta ei saada pingi ette ja lööb tickeri segi. kuigi muu toimib.") */

				return
			}
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	//checkib kas on elus
	/* log.Println("pong") */
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}
