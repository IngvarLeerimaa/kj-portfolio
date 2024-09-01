package backend

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var uid int
var db *sql.DB
var data Webdata

//var data Webdata

func OpenDB() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "./backend/data/database.db")
	if err != nil {
		fmt.Println("Error opening db:", err)
		http.Error(nil, "Error opening db", http.StatusInternalServerError)
	}

	return db
}

func CreateTables() {
	queries, err := os.ReadFile("./backend/data/tables.sql")
	if err != nil {
		fmt.Println("Error reading tables.sql:", err)
		http.Error(nil, "Error reading queries", http.StatusInternalServerError)
	}
	_, err = db.Exec(string(queries))
	fmt.Println("Creating tables")
	if err != nil {
		fmt.Println("Error creating tables:", err)
		http.Error(nil, "Error creating tables", http.StatusInternalServerError)
	}
}

func CheckLogin(u User) bool {
	username, err := db.Query(`SELECT "userID", "password" FROM "users" WHERE "username" = ?`, u.Username)
	if err != nil {
		fmt.Println("Error checking login:", err)
		http.Error(nil, "Error checking login", http.StatusInternalServerError)
	}
	defer username.Close()
	var pwdHash string
	if username.Next() {
		username.Scan(&uid, &pwdHash)
		err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(u.Password))
		return err == nil
	}
	return false
}

func AddSession(s string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "sessionID",
		Value:  s,
		MaxAge: 3600,
	})

	_, err := db.Exec(`DELETE FROM "sessions" WHERE "userID" = ? OR "timestamp" < DATETIME('NOW', '-1 HOUR')`, uid)
	if err != nil {
		fmt.Println("Error deleting session:", err)
		http.Error(nil, "Error deleting session", http.StatusInternalServerError)
	}

	_, err = db.Exec(`INSERT INTO "sessions"("sessionID", "userID") VALUES(?, ?)`, s, uid)
	if err != nil {
		fmt.Println("Error adding session:", err)
		http.Error(nil, "Error adding session", http.StatusInternalServerError)
	}
}

func CheckSession(c string) (int, string, bool) {
	session, err := db.Query(`SELECT "userID" FROM "sessions" WHERE "sessionID" = ?`, c)
	if err != nil {
		fmt.Println("Error checking session:", err)
		http.Error(nil, "Error checking session", http.StatusInternalServerError)
	}
	defer session.Close()
	//???????
	id := 0
	username := ""
	if session.Next() {
		session.Scan(&id)
		user, err := db.Query(`SELECT username FROM "users" WHERE "userID" = ?`, id)
		if err != nil {
			fmt.Println("Error checking session:", err)
			http.Error(nil, "Error checking session", http.StatusInternalServerError)
		}
		user.Next()
		user.Scan(&username)
		return id, username, true
	}
	return id, username, false
}

func (wd *Webdata) AddCategories() {
	tmpCat := Category{}
	tmpCategories := []Category{}
	categories, err := db.Query(`SELECT * FROM "categories"`)
	if err != nil {
		fmt.Println("Error getting categories:", err)
		http.Error(nil, "Error getting categories", http.StatusInternalServerError)
	}
	defer categories.Close()

	for categories.Next() {
		categories.Scan(&tmpCat.Id, &tmpCat.Name, &tmpCat.Description)
		tmpCategories = append(tmpCategories, tmpCat)
	}
	wd.Categories = tmpCategories
}

// Lisan otp kontroll kui kunagi vaja
func IsLoggedIn(r *http.Request) (int, string, bool) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return 0, "", false
	}
	return CheckSession(cookie.Value)
}

func CheckThread(tId int) bool {
	thread, err := db.Query(`SELECT * FROM threads WHERE threadID = ?`, tId)
	if err != nil {
		fmt.Println("Error checking thread:", err)
		http.Error(nil, "Error checking thread", http.StatusInternalServerError)
	}
	defer thread.Close()
	return thread.Next()
}

// eem. tuimalt kopisin ei tea kas siin on vaja midagi muuta
func (wd *Webdata) FilterThreads(id int) {
	stmt, err := db.Prepare(`SELECT * FROM threads WHERE userID = ?`)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		http.Error(nil, "Error preparing statement", http.StatusInternalServerError)
	}
	threads, err := stmt.Query(id)
	if err != nil {
		fmt.Println("Error getting threads for filter:", err)
		http.Error(nil, "Error getting threads for filter:", http.StatusInternalServerError)
	}
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
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		http.Error(nil, "Error preparing statement", http.StatusInternalServerError)
	}
	likedIds, err := stmt.Query(id)
	if err != nil {
		fmt.Println("Error getting liked threads:", err)
		http.Error(nil, "Error getting liked threads", http.StatusInternalServerError)
	}
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

// CheckErr panics if err is not nil. Just a helper function to reduce code duplication.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func AddMessage(from, to int, message string) {
	user, err := db.Query(`SELECT "username" FROM "users" WHERE "userID" = ?`, from)
	if err != nil {
		fmt.Println("Error getting username:", err)
		http.Error(nil, "Error getting username", http.StatusInternalServerError)
	}
	username := ""
	if user.Next() {
		user.Scan(&username)
	}
	user.Close()
	_, err = db.Exec(`INSERT INTO "chat"("fromID", "toID", "message", fromUser) VALUES(?, ?, ?, ?)`, from, to, message, username)
	if err != nil {
		fmt.Println("Error adding message:", err)
		http.Error(nil, "Error adding message", http.StatusInternalServerError)
	}
}

func GetMessages(from, to, offset int) []Message {
	var messages []Message
	var message Message
	msgs, err := db.Query(`SELECT fromUser, message, timestamp FROM "chat" WHERE ("fromID" = ? AND "toID" = ?) OR ("fromID" = ? AND "toID" = ?) ORDER BY rowid DESC LIMIT 10 OFFSET ?`, from, to, to, from, offset)
	if err != nil {
		fmt.Println("Error getting messages:", err)
		http.Error(nil, "Error getting messages", http.StatusInternalServerError)
	}
	defer msgs.Close()

	for msgs.Next() {
		msgs.Scan(&message.From, &message.Content, &message.Created)
		messages = append(messages, message)
	}
	return messages
}

func GetLastMessageDate(from, to int) (time.Time, bool) {
	var lastDate time.Time
	date, err := db.Query(`SELECT timestamp FROM "chat" WHERE ("fromID" = ? AND "toID" = ?) OR ("fromID" = ? AND "toID" = ?) ORDER BY rowid DESC LIMIT 1`, from, to, to, from)
	if err != nil {
		fmt.Println("Error getting date:", err)
		http.Error(nil, "Error getting date", http.StatusInternalServerError)
	}
	defer date.Close()

	if date.Next() {
		date.Scan(&lastDate)
		return lastDate, true
	}
	return lastDate, false
}
