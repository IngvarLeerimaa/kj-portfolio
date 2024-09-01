package main

import (
	"fmt"
	"net/http"
)

// CheckErr serves 500 if err is not nil. Just a helper function to reduce code duplication.
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		http.Error(nil, "Internal server error", http.StatusInternalServerError)
	}
}

// isLoggedIn takes an HTTP request and checks if the user is logged in by checking if the session cookie exists in the database. If the cookie exists, it returns the user ID and true, otherwise it returns 0 and false.
func isLoggedIn(r *http.Request) (int, bool) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return 0, false
	}
	return CheckSession(cookie.Value)
}

// CheckThread takes a thread ID and checks if the thread exists in the database
func CheckThread(tId int) bool {
	thread, err := db.Query(`SELECT * FROM threads WHERE threadID = ?`, tId)
	CheckErr(err)
	defer thread.Close()
	return thread.Next()
}

// FilterThreads takes a user ID and retrieves all threads created by that user, as well as all threads that the user has liked. It prepares and executes two SQL statements to retrieve the relevant threads from the database. It then populates the Webdata struct with the retrieved threads.
func (wd *Webdata) FilterThreads(id int) {
	stmt, err := db.Prepare(`SELECT * FROM threads WHERE userID = ?`)
	CheckErr(err)
	threads, err := stmt.Query(id)
	CheckErr(err)
	defer threads.Close()
	var tmpThread Thread
	var tmpThreads []Thread

	for threads.Next() {
		threads.Scan(&tmpThread.Id, &tmpThread.User, &tmpThread.Title, &tmpThread.Content, &tmpThread.Created)
		db.QueryRow(`SELECT username FROM users WHERE userID = ?`, tmpThread.User).Scan(&tmpThread.Username)
		db.QueryRow(`SELECT COUNT(*) FROM likes WHERE threadID = ?`, tmpThread.Id).Scan(&tmpThread.Likes)
		db.QueryRow(`SELECT COUNT(*) FROM dislikes WHERE threadID = ?`, tmpThread.Id).Scan(&tmpThread.Dislikes)
		tmpThreads = append(tmpThreads, tmpThread)
	}
	wd.Threads = tmpThreads

	tmpThreads = []Thread{}
	stmt, err = db.Prepare(`SELECT "threadID" FROM likes WHERE userID = ? AND threadID > 0`)
	CheckErr(err)
	likedIds, err := stmt.Query(id)
	CheckErr(err)
	defer likedIds.Close()
	var lId int

	for likedIds.Next() {
		likedIds.Scan(&lId)
		db.QueryRow(`SELECT * FROM "threads" WHERE threadID = ?`, lId).Scan(&tmpThread.Id,
			&tmpThread.User, &tmpThread.Title, &tmpThread.Content, &tmpThread.Created)
		db.QueryRow(`SELECT username FROM users WHERE userID = ?`, tmpThread.User).Scan(&tmpThread.Username)
		db.QueryRow(`SELECT COUNT(*) FROM likes WHERE threadID = ?`, tmpThread.Id).Scan(&tmpThread.Likes)
		db.QueryRow(`SELECT COUNT(*) FROM dislikes WHERE threadID = ?`, tmpThread.Id).Scan(&tmpThread.Dislikes)
		tmpThreads = append(tmpThreads, tmpThread)
	}
	wd.LikedThreads = tmpThreads
}
