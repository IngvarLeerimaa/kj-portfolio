package backend

import (
	"context"
	"fmt"
	"net/http"
)

var ctx context.Context

func SetCtx(ctx1 context.Context) {
	ctx = ctx1
}

func setupHandlers(mux *http.ServeMux) {

	manager := NewManager(ctx)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fileServer := http.FileServer(http.Dir("./frontend/"))

	// Serve static files from the "frontend" directory
	mux.Handle("/", fileServer)
	mux.HandleFunc("/ws", manager.serveWs)
	mux.HandleFunc("/login", manager.loginHandler)
	mux.HandleFunc("/session", manager.sessionHandler)
	mux.Handle("/register", http.HandlerFunc(RegisterHandler))
	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, len(manager.clients))
	})
}

/* func registerHandler(w http.ResponseWriter, r *http.Request) {
	type userRegisterRequest struct {
		Id        int
		Username  string
		FirstName string
		LastName  string
		Age       string
		Gender    string
		Email     string
		Password  string
		Created   time.Time
	}
	var req userRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("ERROR:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Register Body:", r.Body)
	fmt.Println("Register request:", req)
	fmt.Println("Register request:", req.Username)
	type response struct {
		Success  bool   `json:"success"`
		Message  string `json:"message"`
		Redirect string `json:"redirect"`
	}
	resp := response{
		Success:  true,
		Message:  "User registered successfully",
		Redirect: "/login",
	}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
*/
