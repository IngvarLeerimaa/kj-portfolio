package database

import (
	"database/sql"
	"fmt"
	"social-network/pkg/helpers"
)

func SavePost(db *sql.DB, p helpers.Post, groupID int) error {
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	postStmt, err := tx.Prepare("INSERT INTO posts(user_id, privacy, text, image) VALUES(?,?,?,?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer postStmt.Close()

	result, err := postStmt.Exec(p.UserID, p.Privacy, p.Text, p.Image)
	if err != nil {
		tx.Rollback()
		return err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	if p.Privacy == 2 {
		followersStmt, err := tx.Prepare("INSERT INTO specific_followers(post_id, user_id) VALUES(?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer followersStmt.Close()

		for i := 0; i < len(p.Followers); i++ {
			_, err := followersStmt.Exec(postID, p.Followers[i])
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if p.Privacy == 3 {
		groupStmt, err := tx.Prepare("INSERT INTO group_post(post_id, group_id) VALUES(?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer groupStmt.Close()

		_, err = groupStmt.Exec(postID, groupID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func SaveComment(db *sql.DB, c helpers.Comment) error {
	stmt, err := db.Prepare("INSERT INTO comments(post_id, user_id, text, image) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.PostID, c.UserID, c.Text, c.Image)
	if err != nil {
		return err
	}
	return nil
}

func Posts(db *sql.DB, userID, offset int) ([]helpers.Post, error) {
	posts := []helpers.Post{}

	stmt, err := db.Prepare(`SELECT * FROM posts WHERE privacy = 0 OR (user_id = ? AND privacy != 3)
	UNION SELECT * FROM posts WHERE privacy = 1 AND user_id = (SELECT follow_id from follows WHERE user_id = ? AND confirmed = 1)
	UNION SELECT * FROM posts WHERE privacy = 2 AND post_id = (Select post_id from specific_followers WHERE user_id = ?)
	ORDER BY timestamp DESC LIMIT 10 OFFSET ?`)
	if err != nil {
		return posts, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, userID, userID, offset)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		post := helpers.Post{}
		rows.Scan(&post.PostID, &post.UserID, &post.Privacy, &post.Text, &post.Image, &post.Created)
		if post.Image != "" {
			post.Image = "http://localhost:3000/images/" + post.Image
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func UserPosts(db *sql.DB, userID, currentUser int) ([]helpers.Post, error) {
	posts := []helpers.Post{}

	stmt, err := db.Prepare(`SELECT * FROM posts WHERE privacy = 0 AND user_id = ?
	UNION SELECT * FROM posts WHERE privacy = 1 AND user_id = ?
	UNION SELECT * FROM posts WHERE privacy = 2 AND user_id = ? AND post_id = (Select post_id from specific_followers WHERE user_id = ?)
	ORDER BY timestamp DESC`)
	if err != nil {
		return posts, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, userID, userID, currentUser)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		post := helpers.Post{}
		rows.Scan(&post.PostID, &post.UserID, &post.Privacy, &post.Text, &post.Image, &post.Created)
		if post.Image != "" {
			post.Image = "http://localhost:3000/images/" + post.Image
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GroupPosts(db *sql.DB, groupID int) ([]helpers.Post, error) {
	posts := []helpers.Post{}

	stmt, err := db.Prepare(`SELECT * FROM posts WHERE privacy = 3 AND post_id IN (SELECT post_id FROM group_post WHERE group_id = ?)`)
	if err != nil {
		return posts, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(groupID)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		post := helpers.Post{}
		rows.Scan(&post.PostID, &post.UserID, &post.Privacy, &post.Text, &post.Image, &post.Created)
		if post.Image != "" {
			post.Image = "http://localhost:3000/images/" + post.Image
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func Comments(db *sql.DB, userID, postID int) ([]helpers.Comment, error) {
	comments := []helpers.Comment{}

	stmt, err := db.Prepare(`SELECT * FROM posts WHERE post_id = ? AND ((privacy = 0 OR user_id = ?) OR
	(privacy = 1 AND user_id = (SELECT follow_id FROM follows WHERE user_id = ? AND confirmed = 1)) OR
	(privacy = 2 AND post_id = (SELECT post_id FROM specific_followers WHERE user_id = ?)) OR
	(privacy = 3 AND EXISTS (SELECT * FROM group_user WHERE confirmed = 2 AND user_id = ? AND group_id = (SELECT group_id FROM group_post WHERE post_id = ?))))`)
	if err != nil {
		return comments, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(postID, userID, userID, userID, userID, postID)
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	if !rows.Next() {
		return comments, fmt.Errorf("post not found")
	}

	commentStmt, err := db.Prepare(`SELECT * FROM comments where post_id = ?`)
	if err != nil {
		return comments, err
	}
	defer commentStmt.Close()

	commentRows, err := commentStmt.Query(postID)
	if err != nil {
		return comments, err
	}
	defer commentRows.Close()

	for commentRows.Next() {
		c := helpers.Comment{}
		commentRows.Scan(&c.CommentID, &c.PostID, &c.UserID, &c.Text, &c.Image, &c.Created)
		if c.Image != "" {
			c.Image = "http://localhost:3000/images/" + c.Image
		}
		comments = append(comments, c)
	}

	return comments, nil
}
