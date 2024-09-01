package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var input string
var font string
var portNum string
var asciiArt string

func main() {

	validport()

	server(portNum)

}
func server(portNum string) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/download", downloadHandler)

	// terminal feedback
	fmt.Printf("\nServer is running on http://localhost:" + portNum)
	fmt.Printf("\nCtrl + Left Click on the link above\n")
	fmt.Println("To stop the server press Ctrl + C")
	if err := http.ListenAndServe(":"+portNum, mux); err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	var fileName = "templates/index.html"

	t, err := template.ParseFiles(fileName)

	if err != nil {
		fmt.Println(http.StatusInternalServerError) //500
		return
	}

	if r.URL.Path == "/" && r.Method == http.MethodGet {
		t.Execute(w, nil)
		return
	}

	if r.URL.Path == "/ascii-art" && r.Method == http.MethodPost {

		input = r.FormValue("input")
		font = r.FormValue("font")

		if !checkAscii(r.FormValue("input")) {
			http.ServeFile(w, r, "templates/e400.html") //400
		}
		output := createBanner(input, font)

		asciiArt = output

		t.Execute(w, output)
		return
	}

	http.ServeFile(w, r, "templates/e404.html") // 404
}

func createBanner(input string, font string) string {

	var output string

	if font == "" {
		font = "standard"
	}

	file, err := os.ReadFile("banners/" + font + ".txt")
	if err != nil {
		os.Exit(1)
	}

	ascii := strings.Split(strings.Replace(string(file), "\r\n", "\n", -1), "\n")

	userInput := strings.Split(input, "\n")

	for i, letter := range userInput {
		if letter == "" {
			if i != len(userInput)-1 {
				output = output + "\n"
			}
			continue
		}
		for i := 1; i < 9; i++ { // y=8
			for j := 0; j < len(letter); j++ {
				if letter[j] == 10 || letter[j] == 13 {
					continue
				} else if letter[j] > 126 || letter[j] < 32 {
					err := ("Rune out side of ascii scope. Please write something else.\n")
					return err
				}
				index := int(letter[j]-32)*9 + i
				output = output + ascii[index]
			}
			output = output + "\n"

		}
	}

	return output
}

// Ascii check was originally inside the loop but for error 404 had to move it out
func checkAscii(input string) bool {

	for _, value := range input {
		if value < 32 || value > 126 {
			if value != 10 && value != 13 {
				return false
			}
		}
	}
	return true
}

func validport() {

	fmt.Println("Welcome to ascii-art-web")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt the user for a port number
		fmt.Print("Enter a port number (1-65535) you would like to run the server at: \n")

		// Scan the user's input
		scanner.Scan()
		portNum = scanner.Text()

		// Parse the port number from a string to an integer
		port, err := strconv.Atoi(portNum)
		if err != nil {
			// Handle the error if the port is not a valid number
			fmt.Println("Error:", err)
			continue
		}

		// Check that the port is within a valid range (1-65535)
		if port < 1 || port > 65535 {
			fmt.Println("Error: invalid port number")
			continue
		}

		// Try to bind to the port to check if it is free
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			// Handle the error if the port is already in use
			fmt.Println("Error:", err)
			continue
		}
		ln.Close()
		break
	}

}
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "Ascii-art"
	// Create a temporary file
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		// Handle the error
		log.Fatal(err)
	}
	defer file.Close()
	defer os.Remove(file.Name())

	// Write the text to the file
	_, err = file.WriteString(asciiArt)
	if err != nil {
		// Handle the error
		log.Fatal(err)
	}

	// Set the content type and attachment headers
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName+".txt")

	// Serve the file to the user
	http.ServeContent(w, r, fileName, time.Now(), file)
}
