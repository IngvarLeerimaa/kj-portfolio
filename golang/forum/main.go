package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Open the database and create tables
	db = OpenDatabase()
	defer db.Close()
	CreateTables()

	// Define the HTTP request handlers
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/category", CategoryHandler)
	http.HandleFunc("/thread", ThreadHandler)
	http.HandleFunc("/newCategory", NewHandler)
	http.HandleFunc("/newThread", NewThreadHandler)
	http.HandleFunc("/newComment", NewCommentHandler)
	http.HandleFunc("/like", LikesHandler)
	http.HandleFunc("/dislike", DislikesHandler)
	http.HandleFunc("/profile", ProfileHandler)

	// Serve static files from the "style" directory
	fileServer := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style", fileServer))

	//// Start the server and listen for incoming requests
	fmt.Println("Default port is mapped at http://localhost:8080")
	fmt.Println("Default port for Docker is mapped at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
