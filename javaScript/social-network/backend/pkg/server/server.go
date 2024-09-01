package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"social-network/pkg/handlers"
	"social-network/pkg/middleware"
)



func Start(db *sql.DB) {

	router := http.NewServeMux()

	handlers.AddHandlers(router, db)
	
	routerWithMiddleware := middleware.CORS(middleware.Auth(router.ServeHTTP, db))

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", routerWithMiddleware))
}
