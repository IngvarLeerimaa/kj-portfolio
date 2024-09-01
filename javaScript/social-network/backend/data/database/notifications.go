package database

import (
	"database/sql"
	"social-network/pkg/helpers"
)

func Notifications(db *sql.DB, userID int) ([]helpers.Notification, error) {
	notifications, err := followNotification(db, userID)
	if err != nil {
		return notifications, err
	}

	iNotifications, err := inviteNotification(db, userID)
	if err != nil {
		return notifications, err
	}

	notifications = append(notifications, iNotifications...)

	rNotifications, err := requestNotification(db, userID)
	if err != nil {
		return notifications, err
	}

	notifications = append(notifications, rNotifications...)

	eNotifications, err := eventNotification(db, userID)
	if err != nil {
		return notifications, err
	}

	notifications = append(notifications, eNotifications...)

	return notifications, nil
}

func followNotification(db *sql.DB, userID int) ([]helpers.Notification, error) {
	notifications := []helpers.Notification{}
	followStmt, err := db.Prepare(`SELECT user_id, first_name, last_name, avatar FROM users WHERE user_id IN 
							(SELECT follow_id FROM follows WHERE user_id = ? AND confirmed = 0)`)
	if err != nil {
		return notifications, err
	}
	defer followStmt.Close()

	rows, err := followStmt.Query(userID)
	if err != nil {
		return notifications, err
	}
	defer rows.Close()

	for rows.Next() {
		n := helpers.Notification{}
		rows.Scan(&n.User.ID, &n.User.FirstName, &n.User.LastName, &n.User.Avatar)
		n.User.Avatar = "http://localhost:3000/images/" + n.User.Avatar
		n.NotificationType = "follow"
		notifications = append(notifications, n)
	}

	return notifications, nil
}

func inviteNotification(db *sql.DB, userID int) ([]helpers.Notification, error) {
	notifications := []helpers.Notification{}
	stmt, err := db.Prepare(`SELECT group_id FROM group_user WHERE confirmed = 1 AND user_id = ?`)
	if err != nil {
		return notifications, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return notifications, err
	}
	defer rows.Close()

	for rows.Next() {
		n := helpers.Notification{}
		rows.Scan(&n.GroupID)
		n.NotificationType = "invite"
		notifications = append(notifications, n)
	}

	return notifications, nil
}

func requestNotification(db *sql.DB, userID int) ([]helpers.Notification, error) {
	notifications := []helpers.Notification{}
	stmt, err := db.Prepare(`SELECT group_id, user_id FROM group_user WHERE confirmed = 0 AND group_id IN (SELECT group_id FROM groups WHERE admin_id = ?)`)
	if err != nil {
		return notifications, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return notifications, err
	}
	defer rows.Close()

	for rows.Next() {
		n := helpers.Notification{}
		rows.Scan(&n.GroupID, &n.UserID)
		n.NotificationType = "request"
		notifications = append(notifications, n)
	}

	return notifications, nil
}

func eventNotification(db *sql.DB, userID int) ([]helpers.Notification, error) {
	notifications := []helpers.Notification{}

	tx, err := db.Begin()
	if err != nil {
		return notifications, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`SELECT DISTINCT event_id FROM event_users WHERE event_id IN (SELECT event_id FROM events WHERE group_id IN
		(SELECT group_id FROM group_user WHERE user_id = ? AND confirmed = 2)) AND event_id NOT IN (SELECT event_id FROM event_users WHERE user_id = ?)
		UNION SELECT event_id FROM events WHERE group_id IN (SELECT group_id from group_user WHERE user_id = ?) AND 
		event_id NOT IN (SELECT event_id FROM event_users);`)
	if err != nil {
		return notifications, err
	}
	defer stmt.Close()

	eStmt, err := tx.Prepare(`SELECT * FROM events WHERE event_id = ?`)
	if err != nil {
		return notifications, err
	}
	defer eStmt.Close()

	rows, err := stmt.Query(userID, userID, userID)
	if err != nil {
		return notifications, err
	}
	defer rows.Close()

	for rows.Next() {
		n := helpers.Notification{}
		id := 0
		rows.Scan(&id)
		eRows, err := eStmt.Query(id)
		if err != nil {
			return notifications, err
		}
		defer eRows.Close()
		if eRows.Next() {
			eRows.Scan(&n.GroupEvent.EventID, &n.GroupEvent.GroupID, &n.GroupEvent.CreatorID, &n.GroupEvent.Title, &n.GroupEvent.Description, &n.GroupEvent.Date, &n.GroupEvent.Created)
		}
		n.NotificationType = "event"
		notifications = append(notifications, n)
	}

	if err := tx.Commit(); err != nil {
		return notifications, err
	}

	return notifications, nil
}
