package backend

import (
	"fmt"
	"net/http"
	"strconv"
)

func SetUpServer() {
	mux := http.NewServeMux()

	setupHandlers(mux)

	portNum = strconv.Itoa(portVerifeir())
	port := "https://localhost:" + portNum + "/"
	setPortNum(portNum)

	// Start the HTTP server on the specified port
	fmt.Println("Server started on " + port)
	err := http.ListenAndServeTLS(":"+portNum, "server.crt", "server.key", mux)
	if err != nil {
		panic(err)
	}

}
