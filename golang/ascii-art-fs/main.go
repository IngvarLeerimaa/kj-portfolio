package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	output := createBanner()
	fmt.Print(output)

}

func getInput() string {
	input := os.Args[1]
	return input
}

func checkBanner(bannerName string) bool {
	if bannerName == "standard" || bannerName == "shadow" || bannerName == "thinkertoy" {
		return true
	}
	return false
}

func createBanner() string {

	var output string

	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Printf("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n")
		os.Exit(0)
	}

	bannerName := "standard"
	if len(os.Args) == 3 {
		re := regexp.MustCompile(`(.*)\.txt`)
		bannerName = re.ReplaceAllString(os.Args[2], "$1")
	}

	if !checkBanner(bannerName) {
		fmt.Println("Available banners are: shadow, standard, tinkertoy.\\nEx: go run . something standard.\nPlease try again.")
		os.Exit(1)
	}

	file, err := os.ReadFile("banners/" + bannerName + ".txt")
	if err != nil {
		os.Exit(2)
	}

	ascii := strings.Split(strings.Replace(string(file), "\r\n", "\n", -1), "\n")

	input := getInput()

	userInput := strings.Split(input, "\\n")

	for i, letter := range userInput {
		if letter == "" {
			if i != len(userInput)-1 {
				output = output + "\n"
			}
			continue
		}
		for i := 1; i < 9; i++ { // y=8
			for j := 0; j < len(letter); j++ {
				if letter[j] > 126 || letter[j] < 32 {
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
