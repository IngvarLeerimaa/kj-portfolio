package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	Rooms = make(map[int]*Room)
	AvailableRooms = make(map[int]*Room)
	clientManager := NewManager()
	router.HandleFunc("/ws", clientManager.WebSocketHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}
