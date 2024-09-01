package database

import (
	"database/sql"
	"errors"
	"social-network/pkg/helpers"
)

func SaveGroup(db *sql.DB, group helpers.Group) error {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	groupSt, err := tx.Prepare("INSERT INTO groups(title, description, admin_id) VALUES (?,?,?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer groupSt.Close()

	result, err := groupSt.Exec(group.Title, group.Description, group.AdminID)
	if err != nil {
		tx.Rollback()
		return err
	}

	groupID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	userStmt, err := tx.Prepare("INSERT INTO group_user(group_id, user_id, confirmed) VALUES(?,?, 2)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer userStmt.Close()

	_, err = userStmt.Exec(groupID, group.AdminID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

func Groups(db *sql.DB) ([]helpers.Group, error) {
	groups := []helpers.Group{}
	stmt, err := db.Prepare("SELECT group_id, title, description, admin_id FROM groups")
	if err != nil {
		return groups, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return groups, err
	}
	defer rows.Close()

	for rows.Next() {
		g := helpers.Group{}
		rows.Scan(&g.GroupID, &g.Title, &g.Description, &g.AdminID)
		groups = append(groups, g)
	}

	return groups, nil
}

func JoinData(db *sql.DB, GroupID, currentUser int) (bool, bool, bool, error) {
	j, i, r := false, false, false
	jStmt, err := db.Prepare("SELECT confirmed FROM group_user WHERE group_id = ? AND user_id = ?")
	if err != nil {
		return j, i, r, err
	}
	defer jStmt.Close()

	rows, err := jStmt.Query(GroupID, currentUser)
	if err != nil {
		return j, i, r, err
	}
	defer rows.Close()

	if rows.Next() {
		var confirmed int
		rows.Scan(&confirmed)
		if confirmed == 2 {
			j = true
		} else if confirmed == 1 {
			i = true
		} else {
			r = true
		}
	}

	return j, i, r, nil
}

func IsGroupMember(db *sql.DB, groupID, currentUser int) (bool, error) {
	stmt, err := db.Prepare("SELECT confirmed FROM group_user WHERE group_id = ? AND user_id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID, currentUser)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, nil
	}
	confirmed := 0
	rows.Scan(&confirmed)

	return confirmed == 2, nil
}

func GroupAdminID(db *sql.DB, groupID int) (int, error) {
	adminID := 0
	stmt, err := db.Prepare("SELECT admin_id FROM groups WHERE group_id = ?")
	if err != nil {
		return adminID, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID)
	if err != nil {
		return adminID, err
	}
	defer rows.Close()

	if !rows.Next() {
		return adminID, errors.New("invalid group id")
	}

	rows.Scan(&adminID)

	return adminID, nil
}

func MembersData(db *sql.DB, groupID int) ([]int, []int, error) {
	var joined []int
	var invited []int
	stmt, err := db.Prepare("SELECT user_id, confirmed FROM group_user WHERE group_id = ? AND (confirmed = 2 OR confirmed = 1)")
	if err != nil {
		return joined, invited, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID)
	if err != nil {
		return joined, invited, err
	}
	defer rows.Close()

	for rows.Next() {
		i := 0
		uID := 0
		rows.Scan(&uID, &i)
		if i == 2 {
			joined = append(joined, uID)
		} else if i == 1 {
			invited = append(invited, uID)
		}
	}
	return joined, invited, nil
}

func AddEvent(db *sql.DB, e helpers.GroupEvent) (int64, error) {
	n := int64(0)
	stmt, err := db.Prepare("INSERT INTO events(group_id, user_id, title, description, date) VALUES (?,?,?,?,?)")
	if err != nil {
		return n, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.GroupID, e.CreatorID, e.Title, e.Description, e.Date)
	if err != nil {
		return n, err
	}

	return result.LastInsertId()
}

func GroupEvents(db *sql.DB, groupID, userID int) ([]helpers.GroupEvent, error) {
	events := []helpers.GroupEvent{}

	stmt, err := db.Prepare(`SELECT * FROM events WHERE  group_id = ?`)
	if err != nil {
		return events, err
	}
	defer stmt.Close()

	attendStmt, err := db.Prepare(`SELECT going FROM event_users WHERE  event_id = ? AND user_id = ?`)
	if err != nil {
		return events, err
	}
	defer attendStmt.Close()

	rows, err := stmt.Query(groupID)
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		event := helpers.GroupEvent{}
		rows.Scan(&event.EventID, &event.GroupID, &event.CreatorID, &event.Title, &event.Description, &event.Date, &event.Created)

		aRows, err := attendStmt.Query(event.EventID, userID)
		if err != nil {
			return events, err
		}
		defer aRows.Close()

		if aRows.Next() {
			going := 0
			aRows.Scan(&going)
			if going == 1 {
				event.Going = true
			} else {
				event.NotGoing = true
			}
		}

		events = append(events, event)
	}
	return events, nil
}

func GetGroupID(db *sql.DB, eventID int) (int, error) {
	groupID := 0
	stmt, err := db.Prepare("SELECT group_id from events WHERE event_id = ?")
	if err != nil {
		return groupID, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(eventID)
	if err != nil {
		return groupID, err
	}
	defer rows.Close()

	if !rows.Next() {
		return groupID, errors.New("invalid event id")
	}

	rows.Scan(&groupID)

	return groupID, nil
}

func AddEventUser(db *sql.DB, eventID, userID int, going bool) error {
	stmt, err := db.Prepare("INSERT INTO event_users(event_id, user_id, going) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	g := 0
	if going {
		g = 1
	}

	_, err = stmt.Exec(eventID, userID, g)
	if err != nil {
		return err
	}

	return nil
}

func GroupUsers(db *sql.DB, groupID int) ([]int, error) {
	users := []int{}
	stmt, err := db.Prepare("SELECT user_id FROM group_user WHERE group_id = ?")
	if err != nil {
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		id := 0
		rows.Scan(&id)
		users = append(users, id)
	}

	return users, nil
}

func AddGroupUser(db *sql.DB, groupID, userID, confirmed int) error {
	stmt, err := db.Prepare("INSERT INTO group_user (group_id, user_id, confirmed) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(groupID, userID, confirmed)

	return err
}

func IsPending(db *sql.DB, groupID, userID, confirmed int) (bool, error) {
	stmt, err := db.Prepare("SELECT * from group_user WHERE group_id = ? AND user_id = ? and confirmed = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID, userID, confirmed)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}

func ConfirmGroupUser(db *sql.DB, groupID, userID int) error {
	stmt, err := db.Prepare("UPDATE group_user SET confirmed = 2 WHERE group_id = ? AND user_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(groupID, userID)

	return err
}

func DeleteGroupUser(db *sql.DB, groupID, userID int) error {
	stmt, err := db.Prepare("DELETE FROM group_user WHERE group_id = ? AND user_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(groupID, userID)

	return err
}
