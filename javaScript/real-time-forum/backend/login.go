package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (m *Manager) loginHandler(w http.ResponseWriter, r *http.Request) {
	type userLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req userLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	re := regexp.MustCompile(`^\S+@\S+\.\S+$`)
	var user *sql.Rows
	var err error

	if !re.MatchString(req.Username) {
		user, err = db.Query("SELECT userID, password FROM Users WHERE username = ?", req.Username)
		if err != nil {
			log.Println(err)
			return
		}
		defer user.Close()
	} else {
		user, err = db.Query("SELECT userID, password FROM Users WHERE email = ?", req.Username)
		if err != nil {
			log.Println(err)
			return
		}
		defer user.Close()
	}

	if user.Next() {

		var passwordHash string
		user.Scan(&uid, &passwordHash)
		if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password)) == nil {
			user.Close()
			sessionId := uuid.Must(uuid.NewV4()).String()
			AddSession(sessionId, w)
			type response struct {
				OTP      string `json:"otp"`
				Redirect string `json:"redirect"`
			}

			otp := m.otps.NewOTP()
			resp := response{
				OTP:      otp.Key,
				Redirect: "/",
			}
			data, err := json.Marshal(resp)
			if err != nil {
				log.Println(err)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return
		} else {
			//invalid password
			w.WriteHeader(http.StatusUnauthorized)
		}

	} else {
		// invalid username/email
		w.WriteHeader(http.StatusUnauthorized)
	}

}

/* if req.Username == "asd" && req.Password == "asd" {
	type response struct {
		OTP      string `json:"otp"`
		Redirect string `json:"redirect"`
	}

	otp := m.otps.NewOTP()
	resp := response{
		OTP:      otp.Key,
		Redirect: "/",
	}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

if req.Username == "qwe" && req.Password == "qwe" {
	type response struct {
		OTP      string `json:"otp"`
		Redirect string `json:"redirect"`
	}

	otp := m.otps.NewOTP()
	resp := response{
		OTP:      otp.Key,
		Redirect: "/",
	}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
w.WriteHeader(http.StatusUnauthorized) */
