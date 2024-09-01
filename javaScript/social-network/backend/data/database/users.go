package database

import (
	"database/sql"
	"log"
	"social-network/pkg/helpers"
)

func EmailExists(db *sql.DB, email string) (string, error) {
	stmt, err := db.Prepare("SELECT password FROM users WHERE email = ?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	rows, err := stmt.Query(email)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	passHash := ""
	if rows.Next() {
		rows.Scan(&passHash)
	}
	return passHash, nil
}

func SaveUserDetails(db *sql.DB, u helpers.User) error {
	stmt, err := db.Prepare("INSERT INTO users(email, password, first_name, last_name, date_of_birth, avatar, nickname, about) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Email, u.Password, u.FirstName, u.LastName, u.DateOfBirth, u.Avatar, u.Nickname, u.About)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *sql.DB, userID, privacy int) error {
	stmt, err := db.Prepare("UPDATE users SET public = ? WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(privacy, userID)
	if err != nil {
		return err
	}
	return nil
}

func SaveSession(db *sql.DB, sessionID, email string) (int, error) {
	qStmt, err := db.Prepare("SELECT user_id FROM users WHERE email = ?")
	if err != nil {
		return 0, err
	}

	rows, err := qStmt.Query(email)
	if err != nil {
		return 0, err
	}
	qStmt.Close()
	userID := 0
	if rows.Next() {
		rows.Scan(&userID)
	}
	rows.Close()
	iStmt, err := db.Prepare("INSERT INTO sessions(session_id, user_id) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}
	defer iStmt.Close()

	_, err = iStmt.Exec(sessionID, userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func DeleteSession(db *sql.DB, sessionID string) {
	stmt, err := db.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionID)
	if err != nil {
		log.Println(err)
	}
}

func ValidateSessionID(db *sql.DB, sessionID string) (int, error) {
	dStmt, err := db.Prepare("DELETE FROM sessions WHERE timestamp <= datetime('now', '-1 day')")
	if err != nil {
		return 0, err
	}

	_, err = dStmt.Exec()
	if err != nil {
		return 0, err
	}

	dStmt.Close()

	sStmt, err := db.Prepare("SELECT user_id FROM sessions WHERE session_id = ?")
	if err != nil {
		return 0, err
	}
	defer sStmt.Close()

	rows, err := sStmt.Query(sessionID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	userID := 0
	if rows.Next() {
		rows.Scan(&userID)
	}

	return userID, nil
}

func User(db *sql.DB, userID, currentUser int) (helpers.User, error) {
	user := helpers.User{}

	stmt, err := db.Prepare("SELECT user_id, email, first_name, last_name, date_of_birth, avatar, nickname, about, public FROM users WHERE user_id = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.Avatar, &user.Nickname, &user.About, &user.Public)
		user.Avatar = "http://localhost:3000/images/" + user.Avatar
	}
	user.CurrentUser = userID == currentUser

	return user, nil
}

func UserList(db *sql.DB, currentUser int) ([]helpers.User, error) {
	users := []helpers.User{}
	stmt, err := db.Prepare("SELECT user_id, first_name, last_name, avatar, public FROM users")
	if err != nil {
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user := helpers.User{}
		rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Avatar, &user.Public)
		user.Avatar = "http://localhost:3000/images/" + user.Avatar
		user.CurrentUser = user.ID == currentUser
		users = append(users, user)
	}

	return users, nil
}

func AddFollow(db *sql.DB, userID, currentUser int) error {
	stmt, err := db.Prepare(`INSERT INTO follows (user_id, follow_id, confirmed) SELECT ?, ?,
							CASE WHEN u.public = 1 THEN 1 ELSE 0 END AS confirmed FROM users u WHERE u.user_id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, currentUser, userID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFollow(db *sql.DB, userID, currentUser int) error {
	stmt, err := db.Prepare(`DELETE FROM follows WHERE user_id = ? AND follow_id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, currentUser)
	if err != nil {
		return err
	}

	return nil
}

func ConfirmFollow(db *sql.DB, userID, currentUser int) error {
	stmt, err := db.Prepare("UPDATE follows SET confirmed = 1 WHERE follow_id = ? AND user_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, currentUser)
	if err != nil {
		return err
	}

	return nil
}

func FollowData(db *sql.DB, userID, currentUser int) (bool, bool, bool, error) {
	following, pending, follower := false, false, false
	followingStmt, err := db.Prepare("SELECT confirmed FROM follows WHERE user_id = ? AND follow_id = ?")
	if err != nil {
		return following, pending, follower, err
	}

	rows, err := followingStmt.Query(userID, currentUser)
	if err != nil {
		return following, pending, follower, err
	}
	followingStmt.Close()

	if rows.Next() {
		var confirmed int
		rows.Scan(&confirmed)
		if confirmed == 1 {
			following = true
		} else {
			pending = true
		}
	}
	rows.Close()

	followerStmt, err := db.Prepare("SELECT * FROM follows WHERE user_id = ? AND follow_id = ? AND confirmed = 1")
	if err != nil {
		return following, pending, follower, err
	}
	defer followerStmt.Close()

	newRows, err := followerStmt.Query(currentUser, userID)
	if err != nil {
		return following, pending, follower, err
	}
	defer newRows.Close()

	return following, pending, newRows.Next(), nil
}

func Followers(db *sql.DB, userID int) ([]helpers.User, error) {
	followers := []helpers.User{}
	stmt, err := db.Prepare(`SELECT user_id, first_name, last_name, avatar FROM users WHERE user_id IN 
							(SELECT follow_id FROM follows WHERE user_id = ? and confirmed = 1)`)
	if err != nil {
		return followers, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return followers, err
	}
	defer rows.Close()

	for rows.Next() {
		follower := helpers.User{}
		rows.Scan(&follower.ID, &follower.FirstName, &follower.LastName, &follower.Avatar)
		follower.Avatar = "http://localhost:3000/images/" + follower.Avatar

		followers = append(followers, follower)
	}

	return followers, nil
}

func Following(db *sql.DB, userID int) ([]helpers.User, error) {
	following := []helpers.User{}
	stmt, err := db.Prepare(`SELECT user_id, first_name, last_name, avatar FROM users WHERE user_id IN 
							(SELECT user_id FROM follows WHERE follow_id = ? AND confirmed = 1)`)
	if err != nil {
		return following, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return following, err
	}
	defer rows.Close()

	for rows.Next() {
		follower := helpers.User{}
		rows.Scan(&follower.ID, &follower.FirstName, &follower.LastName, &follower.Avatar)
		follower.Avatar = "http://localhost:3000/images/" + follower.Avatar

		following = append(following, follower)
	}

	return following, nil
}