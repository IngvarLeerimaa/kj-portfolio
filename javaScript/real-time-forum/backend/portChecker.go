package backend

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func portVerifeir() int {
	fmt.Println("--------------------------")

	fmt.Println("Welcome to Real-Time-Forum")
	fmt.Println("Before we get started, I have to ask you for a port number to run the server at...")

	scanner := bufio.NewScanner(os.Stdin)
	var portNum string

	for {
		// Prompt the user for a port number
		fmt.Println("Enter a port number (1-65535) you would like to run the server at")
		fmt.Println("Press enter to use the default port 8080")

		// Scan the user's input
		scanner.Scan()
		portNum = scanner.Text()

		if portNum == "" {
			portNum = "8080"
		}

		// Parse the port number from a string to an integer
		port, err := strconv.Atoi(portNum)
		if err != nil {
			// Handle the error if the port is not a valid number
			fmt.Println("Hmm.. Something went wrong")
			fmt.Println("Error:", err)
			continue
		}

		// Check that the port is within a valid range (1-65535)
		if port < 1 || port > 65535 {
			fmt.Println("Hmm.. Something went wrong:")
			fmt.Println("Error: invalid port number")
			continue
		}

		// Try to bind to the port to check if it is free
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			// Handle the error if the port is already in use
			fmt.Println("Hmm.. Something went wrong:")
			fmt.Println("Error:", err)
			continue
		}
		ln.Close()
		return port
	}
}
