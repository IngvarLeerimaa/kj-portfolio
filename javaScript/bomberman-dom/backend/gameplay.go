package main

import (
	"encoding/json"
	"log"
	"math/rand"
)

type ExplosionCell struct {
	Type string `json:"type"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	ID   int    `json:"id"`
}

func (r *Room) PrepareGame() {
	data := struct {
		Type         string   `json:"type"`
		Grid         []string `json:"grid"`
		PlayerNumber int      `json:"playerNumber"`
	}{}

	grid := []string{}
	rows := 13
	cols := 21

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i%2 == 1 && j%2 == 1 {
				grid = append(grid, "wall")
			} else if !((i <= 1 || i >= rows-2) && (j <= 1 || j >= cols-2)) && rand.Intn(10) <= 2 {
				grid = append(grid, "block")
			} else if i == 0 && j == 0 || i == rows-1 && j == cols-1 || i == 0 && j == cols-1 && len(r.Clients) > 2 || i == rows-1 && j == 0 && len(r.Clients) > 3 {
				grid = append(grid, "player")
			} else {
				grid = append(grid, "")
			}
		}
	}
	data.Type = "prepareGame"
	data.Grid = grid

	for i, client := range r.Clients {
		data.PlayerNumber = i + 1

		msgJSON, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		client.mutex.Lock()
		defer client.mutex.Unlock()
		err = client.connection.WriteMessage(1, msgJSON)
		if err != nil {
			log.Println(err)
		}
	}

}

func (r *Room) SendPlayerLocation(c *Client, x, y int, path string) {
	data := struct {
		Type   string `json:"type"`
		Player int    `json:"player"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		Path string `json:"path"`
	}{}
	data.Player = c.num
	data.Type = "playerLocation"
	data.X = x
	data.Y = y
	data.Path = path

	msgJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	for _, client := range r.Clients {
		if client != c {
			client.mutex.Lock()
			defer client.mutex.Unlock()
			err = client.connection.WriteMessage(1, msgJSON)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (r *Room) SendBomb(x, y int) {
	data := struct {
		Type string `json:"type"`
		X    int    `json:"x"`
		Y    int    `json:"y"`
	}{
		Type: "plantBomb",
		X:    x,
		Y:    y,
	}

	msgJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	for _, client := range r.Clients {
		client.mutex.Lock()
		defer client.mutex.Unlock()
		err = client.connection.WriteMessage(1, msgJSON)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) SendExplosion(cells []ExplosionCell) {
	playerHit := []bool{false, false, false, false}
	explosions := []ExplosionCell{}

	for _, cell := range cells {
		tmpCell := ExplosionCell{}
		tmpCell.X = cell.X
		tmpCell.Y = cell.Y
		switch cell.Type {
		case "block":
			if rand.Intn(3) == 0 {
				tmpCell.Type = "powerup"
				tmpCell.ID = rand.Intn(3)
			} else {
				tmpCell.Type = "empty"
			}
			explosions = append(explosions, tmpCell)
		case "player":
			playerHit[cell.ID-1] = true
		case "bomb":
			tmpCell.Type = "bomb"
			explosions = append(explosions, tmpCell)
		default:
			tmpCell.Type = "empty"
			explosions = append(explosions, tmpCell)
		}
	}

	data := struct {
		Type  string          `json:"type"`
		Cells []ExplosionCell `json:"cells"`
	}{
		Type: "explosions",
	}

	for _, client := range r.Clients {
		tmpExplosions := explosions
		if playerHit[client.num] {
			tmpExplosions = append(tmpExplosions, ExplosionCell{Type: "life"})
		}

		data.Cells = tmpExplosions

		msgJSON, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		client.mutex.Lock()
		defer client.mutex.Unlock()
		err = client.connection.WriteMessage(1, msgJSON)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) SendRemovePowerUp(id int) {
	data := struct {
		Type string `json:"type"`
		ID   int    `json:"id"`
	}{
		Type: "rmPowerUp",
		ID:   id,
	}

	msgJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	for _, client := range r.Clients {
		client.mutex.Lock()
		defer client.mutex.Unlock()
		err = client.connection.WriteMessage(1, msgJSON)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) SendRemovePlayer(id int) {
	data := struct {
		Type string `json:"type"`
		ID   int    `json:"id"`
	}{
		Type: "rmPlayer",
		ID:   id,
	}

	msgJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	for _, client := range r.Clients {
		client.mutex.Lock()
		defer client.mutex.Unlock()
		err = client.connection.WriteMessage(1, msgJSON)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Room) SendWinner() {
	data := struct {
		Type string `json:"type"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		Type: "winner",
		ID:   r.Alive[0],
		Name: r.Clients[r.Alive[0]].name,
	}

	msgJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	if len(r.Clients) > 0 {
		for _, client := range r.Clients {
			client.mutex.Lock()
			defer client.mutex.Unlock()
			err = client.connection.WriteMessage(1, msgJSON)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
