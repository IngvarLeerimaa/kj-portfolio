package main

import "time"

type Webdata struct {
	User         User
	Categories   []Category
	Threads      []Thread
	Comments     []Comment
	LikedThreads []Thread
	CategoryName string
}

type User struct {
	Id       int
	Email    string
	Username string
	Password string
	Created  time.Time
}

type Category struct {
	Id          int
	Name        string
	Description string
}

type Thread struct {
	Id       int
	User     int
	Title    string
	Content  string
	Likes    int
	Dislikes int
	Liked    bool
	Disliked bool
	Username string
	Created  time.Time
}

type Comment struct {
	Id       int
	Thread   int
	Title    string
	Text     string
	Likes    int
	Dislikes int
	User     int
	Username string
	Liked    bool
	Disliked bool
	Time     time.Time
}
