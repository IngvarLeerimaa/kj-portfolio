package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(u User) {
	_, err := db.Exec("INSERT INTO users (username, firstname, lastname, age, gender, password, email) VALUES (?, ?, ?, ?, ?, ?, ?)", u.Username, u.FirstName, u.LastName, u.Age, u.Gender, u.Password, u.Email)
	if err != nil {
		fmt.Println("Error registering user:", err)
		http.Error(nil, "Error registering user", http.StatusInternalServerError)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Success  bool   `json:"success"`
		Message  string `json:"message"`
		Redirect string `json:"redirect"`
	}
	if r.Method == "POST" {
		user := User{
			Created: time.Now(),
		}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			fmt.Println("ERROR:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		/* re := regexp.MustCompile(`^\S+@\S+\.\S+$`) */
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			fmt.Println("Error hashing password:", err)
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
		}
		user.Password = string(passwordHash)
		/* 	ageStr := r.FormValue("age")
		fmt.Println(ageStr)
		fmt.Printf("%T\n", ageStr)
		age, err := strconv.Atoi(string(ageStr))
		if err != nil {
			fmt.Println("Error converting age:", err)
			http.Error(w, "Error converting age", http.StatusInternalServerError)
		} */

		//user.Password = string(passwordHash)
		fmt.Println(user)

		resp := response{
			Success:  true,
			Message:  "",
			Redirect: "/login",
		}

		data, err := json.Marshal(resp)
		if err != nil {
			log.Println(err)
			return
		}
		//vbl topelt registreerimine
		//RegisterUser(user)

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
			//selle loogika peab mõtlema ümber, et kuidas saaks kõik errorid kuvada
			/* templates.ExecuteTemplate(w, "register.html", errorMsg) */
			resp.Success = false

			for i := 0; i < len(errorMsg); i++ {
				resp.Message += errorMsg[i] + "\n"
			}
			resp.Redirect = "/register"

			data, err := json.Marshal(resp)
			if err != nil {
				log.Println(err)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(data)
			fmt.Println("resp failed: ", resp)
		} else {
			RegisterUser(user)
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			fmt.Println("User register sucessful:", resp)
			return
			//templates.ExecuteTemplate(w, "login.html", "Registration successful, please login")
		}
	} else {
		if _, _, ok := IsLoggedIn(r); ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			fmt.Println("User already logged in")
			resp := response{
				Success:  false,
				Message:  "This account is already logged in..",
				Redirect: "/register",
			}

			w.WriteHeader(http.StatusMethodNotAllowed)
			data, err := json.Marshal(resp)
			if err != nil {
				log.Println(err)
				return
			}
			w.Write(data)
			fmt.Println("User is already logged in?")
		} else {
			resp := response{
				Success:  false,
				Message:  "Wrong method used to acess this page.",
				Redirect: "/register",
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			fmt.Println("User already logged in")

			w.WriteHeader(http.StatusMethodNotAllowed)
			data, err := json.Marshal(resp)
			if err != nil {
				log.Println(err)
				return
			}
			w.Write(data)
			fmt.Println("Wrong method used to acess this page.")
			/* templates.ExecuteTemplate(w, "register.html", nil) */
		}
	}
}

func (u *User) UsernameExists() bool {
	username, err := db.Query("SELECT username FROM users WHERE username = ?", u.Username)
	if err != nil {
		fmt.Println("Error checking username:", err)
		http.Error(nil, "Error checking username", http.StatusInternalServerError)
	}
	return username.Next()
}

func (u *User) EmailExists() bool {
	email, err := db.Query(`SELECT * FROM "users" WHERE "email" = ?`, u.Email)
	if err != nil {
		fmt.Println("Error checking email:", err)
		http.Error(nil, "Error checking email", http.StatusInternalServerError)
	}
	defer email.Close()
	return email.Next()
}

/*
// AddUser adds the user to the webdata
func (wd *Webdata) AddUser(r *http.Request) {
	tmpuser := User{}
	if id, ok := isLoggedIn(r); ok {
		user, err := db.Query(`SELECT * FROM "users" WHERE "userID" = ?`, id)
		if err != nil {
			fmt.Println("Error getting user:", err)
			http.Error(nil, "Error getting user", http.StatusInternalServerError)
		}
		defer user.Close()
		if user.Next() {
			user.Scan(&tmpuser.Id, &tmpuser.Username, &tmpuser.Email, &tmpuser.Password, &tmpuser.Created)
		}
	}
	wd.User = tmpuser
}
*/
