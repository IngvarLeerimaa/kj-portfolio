package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Room struct {
	ID        int
	Clients   []*Client
	Available bool
	Alive     []int
	Stop      chan struct{}
	sync.RWMutex
}

var (
	Rooms          map[int]*Room
	AvailableRooms map[int]*Room
)

func CreateRoom() *Room {
	return &Room{
		ID:        rand.Intn(1000000000),
		Available: true,
		Stop:      make(chan struct{}),
	}
}

func (r *Room) AddUser(c *Client) {
	r.RWMutex.Lock()
	defer r.RWMutex.Unlock()

	c.num = len(r.Clients)
	r.Alive = append(r.Alive, c.num)
	r.Clients = append(r.Clients, c)
	r.SendPlayerData()

	if len(r.Clients) == 2 {
		go r.counter(20, "twenty")
	}

	if len(r.Clients) == 4 {
		delete(AvailableRooms, r.ID)
		r.Available = false
		r.Stop <- struct{}{}
		go r.counter(10, "ten")
	}
}

func (r *Room) RemoveUser(c *Client) {
	r.RWMutex.Lock()

	defer r.RWMutex.Unlock()
	tmpClients := []*Client{}
	for _, client := range r.Clients {
		if client != c {
			tmpClients = append(tmpClients, client)
		}
	}
	r.Clients = tmpClients
	if len(r.Clients) == 0 {
		delete(AvailableRooms, r.ID)
		delete(Rooms, r.ID)
	}

	if len(r.Clients) == 1 && r.Available {
		r.Stop <- struct{}{}
	}
	if len(r.Clients) > 0 && r.Available {
		AvailableRooms[r.ID] = r
		r.SendPlayerData()
	}

	if !r.Available {
		r.SendRemovePlayer(c.num)
		r.removeAlive(c.num)
		if len(r.Alive) == 1 {
			defer func() {
				recover()
			}()
			r.SendWinner()
		}
	}
}

func (r *Room) SendPlayerData() error {
	type Player struct {
		Name string `json:"name"`
		Self bool   `json:"self"`
	}

	for _, cl := range r.Clients {
		playerData := []Player{}
		for _, c := range r.Clients {
			tmpPlayer := Player{}
			tmpPlayer.Name = c.name
			if c == cl {
				tmpPlayer.Self = true
			} else {
				tmpPlayer.Self = false
			}
			playerData = append(playerData, tmpPlayer)
		}

		data := struct {
			Type       string   `json:"type"`
			PlayerData []Player `json:"playerData"`
		}{
			Type:       "playerData",
			PlayerData: playerData,
		}
		msgJSON, err := json.Marshal(data)
		if err != nil {
			return err
		}
		err = cl.connection.WriteMessage(1, msgJSON)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Room) counter(count int, countType string) error {
	if countType == "ten" {
		r.PrepareGame()
	}
	for i := count; i >= -1; i-- {
		select {
		case <-r.Stop:
			if countType != "ten" {
				return nil
			}
		default:
			data := struct {
				Type  string `json:"type"`
				Count int    `json:"count"`
			}{
				Type:  countType,
				Count: i,
			}
			msgJSON, err := json.Marshal(data)
			if err != nil {
				return err
			}
			for _, c := range r.Clients {
				err = c.connection.WriteMessage(1, msgJSON)
				if err != nil {
					return err
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
	if countType == "twenty" {
		delete(AvailableRooms, r.ID)
		r.Available = false
		go r.counter(10, "ten")
	}
	return nil
}

func (r *Room) BroadcastMessage(c *Client, message string) {
	playerColour := ""
	for i, client := range r.Clients {
		if client == c {
			switch i {
			case 0:
				playerColour = "white"
			case 1:
				playerColour = "rgb(59, 75, 166)"
			case 2:
				playerColour = "rgb(195, 46, 24)"
			case 3:
				playerColour = "rgb(40, 96, 223)"
			default: // should never happen
				playerColour = "ERROR"
			}
		}
	}

	data := struct {
		Type    string `json:"type"`
		From    string `json:"from"`
		Message string `json:"message"`
		Colour  string `json:"colour"`
	}{
		Type:    "message",
		From:    c.name,
		Message: message,
		Colour:  playerColour,
	}
	msgJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling message: ", err)
		return
	}
	for _, c := range r.Clients {
		err = c.connection.WriteMessage(1, msgJSON)
		if err != nil {
			fmt.Println("Error writing message: ", err)
			return
		}
	}

	fmt.Println("Broadcasting message: ", message)
	fmt.Println("Message from: ", c.name)
	fmt.Println("Room ID: ", r.ID)

}

//white, blue, red, lightblue

func (r *Room) removeAlive(id int) {
	tmpAlive := []int{}
	for _, i := range r.Alive {
		if i != id {
			tmpAlive = append(tmpAlive, i)
		}
	}
	r.Alive = tmpAlive
}
