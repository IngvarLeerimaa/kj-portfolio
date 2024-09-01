package database

import (
	"database/sql"
	"social-network/pkg/helpers"
)

func AddMessage(db *sql.DB, msg helpers.Message) error {
	stmt, err := db.Prepare(`INSERT INTO chat (to_id, from_id, message) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(msg.ToID, msg.FromID, msg.Message)
	if err != nil {
		return err
	}

	return nil
}

func AddGroupMessage(db *sql.DB, msg helpers.Message) error {
	stmt, err := db.Prepare(`INSERT INTO group_chat (group_id, from_id, message) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(msg.ToID, msg.FromID, msg.Message)
	if err != nil {
		return err
	}

	return nil
}

func LastMessage(db *sql.DB, to, from int) (helpers.Message, error) {
	message := helpers.Message{}
	stmt, err := db.Prepare(`SELECT * FROM chat WHERE to_id = ? AND from_id = ? OR to_id = ? AND from_id = ? ORDER BY timestamp DESC LIMIT 1`)
	if err != nil {
		return message, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(to, from, from, to)
	if err != nil {
		return message, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&message.ToID, &message.FromID, &message.Message, &message.Created)
	}

	return message, nil
}

func LastGroupMessage(db *sql.DB, groupID int) (helpers.Message, error) {
	message := helpers.Message{}
	stmt, err := db.Prepare(`SELECT * FROM group_chat WHERE group_id = ? ORDER BY timestamp DESC LIMIT 1`)
	if err != nil {
		return message, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID)
	if err != nil {
		return message, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&message.ToID, &message.FromID, &message.Message, &message.Created)
	}

	return message, nil
}

func Messages(db *sql.DB, to, from, offset int) ([]helpers.Message, error) {
	messages := []helpers.Message{}
	stmt, err := db.Prepare(`SELECT * FROM chat WHERE to_id = ? AND from_id = ? OR to_id = ? AND from_id = ? ORDER BY rowid DESC LIMIT 10 OFFSET ?`)
	if err != nil {
		return messages, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(to, from, from, to, offset)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		m := helpers.Message{}
		rows.Scan(&m.ToID, &m.FromID, &m.Message, &m.Created)
		messages = append(messages, m)
	}

	return messages, nil
}

func GroupMessages(db *sql.DB, groupID, offset int) ([]helpers.Message, error) {
	messages := []helpers.Message{}
	stmt, err := db.Prepare(`SELECT * FROM group_chat WHERE group_id = ? ORDER BY rowid DESC LIMIT 10 OFFSET ?`)
	if err != nil {
		return messages, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID, offset)
	if err != nil {
		return messages, err
	}
	defer rows.Close()

	for rows.Next() {
		m := helpers.Message{}
		rows.Scan(&m.ToID, &m.FromID, &m.Message, &m.Created)
		messages = append(messages, m)
	}

	return messages, nil
}
