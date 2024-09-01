package middleware

import (
	"database/sql"
	"net/http"
	"social-network/data/database"
)

func Auth(handler http.HandlerFunc, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/api/v1/user/login" && r.URL.Path != "/api/v1/user/register" {
			sessionCookie, err := r.Cookie("sessionID")
			if err != nil {
				http.Error(w, "Authorization required", http.StatusUnauthorized)
				return
			}

			sessionID := sessionCookie.Value

			isValid, err := database.ValidateSessionID(db, sessionID)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if isValid < 1 {
				http.Error(w, "Authorization required", http.StatusUnauthorized)
				return
			}
		}

		handler(w, r)
	}
}
