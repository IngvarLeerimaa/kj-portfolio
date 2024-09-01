package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var uid int
var db *sql.DB
var data Webdata

// OpenDatabase opens a connection to the SQLite database
func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "database.db")
	CheckErr(err)
	return db
}

// CreateTables creates the tables in the SQLite database by executing queries from a file
func CreateTables() {
	queries, err := os.ReadFile("data/tables.sql")
	CheckErr(err)
	_, err = db.Exec(string(queries))
	CheckErr(err)
}

// EmailExists checks if the email already exists in the database
func (u *User) EmailExists() bool {
	email, err := db.Query(`SELECT * FROM "users" WHERE "email" = ?`, u.Email)
	CheckErr(err)
	defer email.Close()
	return email.Next()
}

// UsernameExists checks if the username already exists in the database
func (u *User) UsernameExists() bool {
	username, err := db.Query(`SELECT * FROM "users" WHERE "username" = ?`, u.Username)
	CheckErr(err)
	defer username.Close()

	return username.Next()
}

// RegisterUser registers a user by inserting their username, password and email into the database
func RegisterUser(u User) {
	_, err := db.Exec(`INSERT INTO users("username", "password", "email") VALUES(?, ?, ?)`, u.Username, u.Password, u.Email)
	CheckErr(err)
}

// CheckLogin checks if the username and password match the ones in the database
func CheckLogin(u User) bool {
	username, err := db.Query(`SELECT "userID", "password" FROM "users" WHERE "username" = ?`, u.Username)
	CheckErr(err)
	var passwordHash string
	defer username.Close()
	if username.Next() {
		username.Scan(&uid, &passwordHash)
		err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(u.Password))
		return err == nil
	}
	return false
}

// AddSession adds a session to the sessions table and sets a cookie in the HTTP response
func AddSession(s string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "sessionID",
		Value:  s,
		MaxAge: 3600,
	})

	_, err := db.Exec(`DELETE FROM "sessions" WHERE "userID" = ? OR "timestamp" < DATETIME('NOW', '-1 HOUR')`, uid)
	CheckErr(err)

	_, err = db.Exec(`INSERT INTO "sessions"("sessionID", "userID") VALUES(?, ?)`, s, uid)
	CheckErr(err)
}

// CheckSession checks if the session cookie is valid comparing it to sessions tagble in the database
func CheckSession(c string) (int, bool) {
	session, err := db.Query(`SELECT "userID" FROM "sessions" WHERE "sessionID" = ?`, c)
	CheckErr(err)
	defer session.Close()
	id := 0
	if session.Next() {
		session.Scan(&id)
		return id, true
	}
	return id, false
}

// AddUserData retrieves user data from the users table and adds it to a Webdata struct
func (wd *Webdata) AddUserData(r *http.Request) {
	tmpuser := User{}
	if id, ok := isLoggedIn(r); ok {
		user, err := db.Query(`SELECT * FROM "users" WHERE "userID" = ?`, id)
		CheckErr(err)
		defer user.Close()
		if user.Next() {
			user.Scan(&tmpuser.Id, &tmpuser.Username, &tmpuser.Email, &tmpuser.Password, &tmpuser.Created)
		}
	}
	wd.User = tmpuser
}

// AddCategories retrieves category data from the categories table and adds it to a Webdata struct
func (wd *Webdata) AddCategories() {
	tmpCat := Category{}
	tmpCategories := []Category{}
	categories, err := db.Query(`SELECT * FROM "categories"`)
	CheckErr(err)
	defer categories.Close()

	for categories.Next() {
		categories.Scan(&tmpCat.Id, &tmpCat.Name, &tmpCat.Description)
		tmpCategories = append(tmpCategories, tmpCat)
	}
	wd.Categories = tmpCategories
}
