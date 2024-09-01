package main

import (
	"net/http"
	"strconv"
)

// AddThreads adds all threads from a category to the Webdata struct
func (wd *Webdata) AddThreads(c int) {
	tmpThread := Thread{}
	tmpThreads := []Thread{}
	var tId int
	threadIds, err := db.Query(`SELECT "threadID" FROM "threadcategories" WHERE "categoryID" = ?`, c)
	CheckErr(err)
	defer threadIds.Close()

	for threadIds.Next() {
		threadIds.Scan(&tId)
		threads, err := db.Query(`SELECT * FROM "threads" WHERE "threadID" = ?`, tId)
		CheckErr(err)
		defer threads.Close()
		threads.Next()
		threads.Scan(&tmpThread.Id, &tmpThread.User,
			&tmpThread.Title, &tmpThread.Content, &tmpThread.Created)
		db.QueryRow(`SELECT COUNT(*) FROM "likes" WHERE threadID = ?`, tId).Scan(&tmpThread.Likes)
		db.QueryRow(`SELECT COUNT(*) FROM "dislikes" WHERE threadID = ?`, tId).Scan(&tmpThread.Dislikes)
		db.QueryRow(`SELECT username FROM users WHERE userID = ?`, tmpThread.User).Scan(&tmpThread.Username)
		tmpThreads = append(tmpThreads, tmpThread)
	}
	wd.Threads = tmpThreads
}

// AddComments adds all comments from a thread to the Webdata struct
func (wd *Webdata) AddComments(t int) {
	tmpThread := Thread{}
	stmt, err := db.Prepare(`SELECT * FROM "threads" WHERE "threadID" = ?`)
	CheckErr(err)
	stmt.QueryRow(t).Scan(&tmpThread.Id, &tmpThread.User, &tmpThread.Title, &tmpThread.Content, &tmpThread.Created)
	db.QueryRow(`SELECT username FROM users WHERE userID = ?`, tmpThread.User).Scan(&tmpThread.Username)
	db.QueryRow(`SELECT COUNT(*) FROM likes WHERE threadID = ?`, tmpThread.Id).Scan(&tmpThread.Likes)
	db.QueryRow(`SELECT COUNT(*) FROM dislikes WHERE threadID = ?`, tmpThread.Id).Scan(&tmpThread.Dislikes)
	db.QueryRow(`SELECT COUNT(*) FROM likes WHERE threadID = ? AND userID = ?`, tmpThread.Id, data.User.Id).Scan(&tmpThread.Liked)
	db.QueryRow(`SELECT COUNT(*) FROM dislikes WHERE threadID = ? AND userID = ?`, tmpThread.Id, data.User.Id).Scan(&tmpThread.Disliked)
	wd.Threads = []Thread{tmpThread}
	tmpComment := Comment{}
	tmpComments := []Comment{}
	threads, err := db.Query(`SELECT * FROM "comments" WHERE "threadID" = ?`, t)
	CheckErr(err)
	defer threads.Close()

	for threads.Next() {
		threads.Scan(&tmpComment.Id, &tmpComment.Thread, &tmpComment.User, &tmpComment.Text, &tmpComment.Time, &tmpComment.Title)
		db.QueryRow(`SELECT username FROM users WHERE userID = ?`, tmpComment.User).Scan(&tmpComment.Username)
		db.QueryRow(`SELECT COUNT(*) FROM likes WHERE commentID = ?`, tmpComment.Id).Scan(&tmpComment.Likes)
		db.QueryRow(`SELECT COUNT(*) FROM dislikes WHERE commentID = ?`, tmpComment.Id).Scan(&tmpComment.Dislikes)
		db.QueryRow(`SELECT COUNT(*) FROM likes WHERE commentID = ? AND userID = ?`, tmpComment.Id, data.User.Id).Scan(&tmpComment.Liked)
		db.QueryRow(`SELECT COUNT(*) FROM dislikes WHERE commentID = ? AND userID = ?`, tmpComment.Id, data.User.Id).Scan(&tmpComment.Disliked)
		tmpComments = append(tmpComments, tmpComment)
	}

	wd.Comments = tmpComments
}

// CreateThread creates a new thread in the database
func CreateThread(r *http.Request) int64 {
	if id, ok := isLoggedIn(r); ok {
		// add thread to database
		stmt, err := db.Prepare(`INSERT INTO "threads"("userID", "title", "content") VALUES(?, ?, ?)`)
		CheckErr(err)
		result, err := stmt.Exec(id, r.FormValue("title"), r.FormValue("content"))
		CheckErr(err)

		// get thread id
		threadId, err := result.LastInsertId()
		CheckErr(err)

		// get selected categories and connect the thread with them
		r.ParseForm()
		for _, v := range r.Form["category"] {
			categoryId, _ := strconv.Atoi(v)
			_, err := db.Exec(`INSERT INTO "threadcategories"("threadID", "categoryID") VALUES(?, ?)`, threadId, categoryId)
			CheckErr(err)
		}
		return threadId
	}
	return 0
}

// AddComment adds a comment to the database
func AddComment(tId, uId int, title, comment string) {
	_, err := db.Exec(`INSERT INTO "comments"("threadID", "userID", "title", "comment") VALUES(?, ?, ?, ?)`, tId, uId, title, comment)
	CheckErr(err)
}

func addLike(cId, uId int, thread string) {
	if thread == "true" {
		stmt, err := db.Prepare(`INSERT OR IGNORE INTO "likes"("threadID", "userID") VALUES (?, ?)`)
		CheckErr(err)
		stmt.Exec(cId, uId)
	} else {
		stmt, err := db.Prepare(`INSERT OR IGNORE INTO "likes"("commentID", "userID") VALUES (?, ?)`)
		CheckErr(err)
		stmt.Exec(cId, uId)
	}
}

// removeLike removes a like from the database
func removeLike(cId, uId int, thread string) {
	if thread == "true" {
		stmt, err := db.Prepare(`DELETE FROM "likes" WHERE "threadID" = ? AND "userID" = ?`)
		CheckErr(err)
		stmt.Exec(cId, uId)
	} else {
		stmt, err := db.Prepare(`DELETE FROM "likes" WHERE "commentID" = ? AND "userID" = ?`)
		CheckErr(err)
		stmt.Exec(cId, uId)
	}
}

// addDislike adds dislike to the database
func addDislike(cId, uId int, thread string) {
	if thread == "true" {
		stmt, err := db.Prepare(`INSERT OR IGNORE INTO "dislikes"("threadID", "userID") VALUES (?, ?)`)
		CheckErr(err)
		stmt.Exec(cId, uId)
	} else {
		stmt, err := db.Prepare(`INSERT OR IGNORE INTO "dislikes"("commentID", "userID") VALUES (?, ?)`)
		CheckErr(err)
		stmt.Exec(cId, uId)
	}
}

// removeDislike removes a dislike from the database
func removeDislike(cId, uId int, thread string) {
	if thread == "true" {
		stmt, err := db.Prepare(`DELETE FROM "dislikes" WHERE "threadID" = ? AND "userID" = ?`)
		CheckErr(err)
		stmt.Exec(cId, uId)
	} else {
		stmt, err := db.Prepare(`DELETE FROM "dislikes" WHERE "commentID" = ? AND "userID" = ?`)
		CheckErr(err)
		stmt.Exec(cId, uId)
	}
}
