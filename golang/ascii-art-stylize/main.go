package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var input string
var font string

func main() {

	server()

}
func server() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// terminal feedback
	fmt.Println("Server is running on http://localhost:8080")
	fmt.Printf("\nCtrl + Left Click on the link above\n")
	fmt.Println("To stop the server press Ctrl + C")
	if err := http.ListenAndServe(":8080", mux); err != nil {
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
