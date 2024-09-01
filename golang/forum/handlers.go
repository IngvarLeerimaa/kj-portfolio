package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// IndexHandler handles the index page, by adding user data and categories to the data, and executing the index.html template.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data.AddUserData(r)
	data.AddCategories()
	templates.ExecuteTemplate(w, "index.html", data)
}

// LoginHandler handles user login, by checking the method of the request, and checking the user's credentials by calling the CheckLogin function. If the user is authenticated, a new session is created and the user is redirected to the index page. Otherwise, an error message is displayed.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		if CheckLogin(user) {
			sessionId := uuid.Must(uuid.NewV4()).String()
			AddSession(sessionId, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			templates.ExecuteTemplate(w, "login.html", "Invalid username or password")
		}
	} else {
		if _, ok := isLoggedIn(r); ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			templates.ExecuteTemplate(w, "login.html", nil)
		}
	}
}

// LogoutHandler logs out the user by deleting the session cookie and redirecting them to the index page.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "sessionID",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// RegisterHandler handles user registration, by checking the method of the request, and creating a new user by calling the RegisterUser function. If the user already exists, an error message is displayed.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		re := regexp.MustCompile(`^\S+@\S+\.\S+$`)
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 10)
		user := User{
			Email:    re.FindString(r.FormValue("email")),
			Username: r.FormValue("username"),
			Password: string(passwordHash),
		}

		var errorMsg []string
		if err != nil {
			errorMsg = append(errorMsg, err.Error())
		}
		if len(user.Email) < 1 || user.EmailExists() {
			errorMsg = append(errorMsg, "Invalid Email address/address already in use")
		}
		if user.UsernameExists() {
			errorMsg = append(errorMsg, "Username already in use")
		}
		if len(errorMsg) > 0 {
			templates.ExecuteTemplate(w, "register.html", errorMsg)
		} else {
			RegisterUser(user)
			templates.ExecuteTemplate(w, "login.html", "Registration successful, please login")
		}
	} else {
		if _, ok := isLoggedIn(r); ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			templates.ExecuteTemplate(w, "register.html", nil)
		}
	}
}

// CategoryHandler handles the category page, by adding user data and categories to the data, and adding threads to the data based on the category ID. If the category does not exist, a 404 error is displayed.
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	cId, _ := strconv.Atoi(r.FormValue("c"))

	data.AddUserData(r)
	data.AddCategories()
	found := false
	for _, v := range data.Categories {
		if v.Id == cId {
			data.AddThreads(cId)
			found = true
		}
	}

	if !found {
		http.NotFound(w, r)
	} else {
		data.CategoryName = data.Categories[cId-1].Name
		templates.ExecuteTemplate(w, "category.html", data)
	}
}

// ThreadHandler handles the thread page, by adding user data and categories to the data, and adding comments to the data based on the thread ID. If the thread does not exist, a 404 error is displayed.
func ThreadHandler(w http.ResponseWriter, r *http.Request) {
	tId, _ := strconv.Atoi(r.FormValue("t"))
	if !CheckThread(tId) {
		http.NotFound(w, r)
		return
	}
	data.AddUserData(r)
	data.AddComments(tId)
	templates.ExecuteTemplate(w, "thread.html", data)
}

// NewHandler handles the new page, by checking the method of the request, and creating a new thread by calling the CreateThread function. If the user is not logged in, an error message is displayed.
func NewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, userID := isLoggedIn(r)
		if !userID {
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}

	}
}

// NewThreadHandler handles the new thread page, by checking the method of the request, and creating a new thread by calling the CreateThread function. If the user is not logged in, they are redirected to the login page.
func NewThreadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if tId := CreateThread(r); tId > 0 {
			http.Redirect(w, r, "/thread?t="+fmt.Sprint(tId), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
	if _, ok := isLoggedIn(r); ok {
		data.AddCategories()
		templates.ExecuteTemplate(w, "newThread.html", data)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// NewCommentHandler handles the new comment page, by checking the method of the request, and creating a new comment by calling the AddComment function. If the user is not logged in, they are redirected to the login page.
func NewCommentHandler(w http.ResponseWriter, r *http.Request) {
	if id, ok := isLoggedIn(r); ok {
		if len(r.FormValue("t")) < 1 {
			http.NotFound(w, r)
			return
		}
		tId, err := strconv.Atoi(r.FormValue("t"))
		CheckErr(err)
		if r.Method == "POST" && ok {
			AddComment(tId, id, r.FormValue("title"), r.FormValue("content"))
			http.Redirect(w, r, "/thread?t="+fmt.Sprint(tId), http.StatusSeeOther)
			return
		}
		templates.ExecuteTemplate(w, "newComment.html", tId)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// LikesHandler handles the likes and dislikes, by checking the method of the request, and adding or removing a like or dislike by calling the addLike, removeLike, addDislike, or removeDislike functions. If the user is not logged in, they are redirected to the login page.
func LikesHandler(w http.ResponseWriter, r *http.Request) {
	cId, _ := strconv.Atoi(r.FormValue("c"))
	if id, ok := isLoggedIn(r); ok {
		if r.FormValue("l") == "add" {
			addLike(cId, id, r.FormValue("t"))
		}
		if r.FormValue("l") == "remove" {
			removeLike(cId, id, r.FormValue("t"))
		}
	}
}

// DislikesHandler handles the likes and dislikes, by checking the method of the request, and adding or removing a like or dislike by calling the addLike, removeLike, addDislike, or removeDislike functions. If the user is not logged in, they are redirected to the login page.
func DislikesHandler(w http.ResponseWriter, r *http.Request) {
	cId, _ := strconv.Atoi(r.FormValue("c"))
	if id, ok := isLoggedIn(r); ok {
		if r.FormValue("l") == "add" {
			addDislike(cId, id, r.FormValue("t"))
		}
		if r.FormValue("l") == "remove" {
			removeDislike(cId, id, r.FormValue("t"))
		}
	}
}

// ProfileHandler handles the profile page, by adding user data and categories to the data, and adding threads to the data based on the user ID. If the user is not logged in, they are redirected to the login page.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if id, ok := isLoggedIn(r); ok {
		data.AddUserData(r)
		data.FilterThreads(id)
		templates.ExecuteTemplate(w, "profile.html", data)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
