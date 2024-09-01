package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	outputtxt, opcount := outputFlag()
	output := createBanner(opcount)

	if opcount == 1 {
		writeFile(output, outputtxt)
	} else {
		fmt.Print(output)
	}
}

func getInput(args []string) string {
	input := args[0]
	return input
}

func checkBanner(bannerName string) bool {
	if bannerName == "standard" || bannerName == "shadow" || bannerName == "thinkertoy" {
		return true
	}
	return false
}

func createBanner(flagcount int) string {

	var output string
	args := os.Args[1+flagcount:]
	if len(args) < 1 || len(args) > 2 {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
		os.Exit(0)
	}

	bannerName := "standard"
	if len(args) == 2 {
		re := regexp.MustCompile(`(.*)\.txt`)
		bannerName = re.ReplaceAllString(args[1], "$1")
	}

	if !checkBanner(bannerName) {
		fmt.Println("Available banners are: shadow, standard, thinkertoy.\n\nEX: go run . --output=<fileName.txt> something standard\nPlease try again.")
		os.Exit(1)
	}

	file, err := os.ReadFile("banners/" + bannerName + ".txt")
	if err != nil {
		os.Exit(2)
	}

	ascii := strings.Split(strings.Replace(string(file), "\r\n", "\n", -1), "\n")

	input := getInput(args)

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

func outputFlag() (string, int) {
	outputptr := flag.String("output", "", "--output=filename.txt")
	flag.Parse()
	oplen := len(*outputptr)
	if oplen < 5 || (*outputptr)[oplen-4:] != ".txt" {
		return "", 0
	}
	return *outputptr, 1
}

func writeFile(output, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		os.Exit(3)
	}
	defer f.Close()
	_, err = f.WriteString(output)
	if err != nil {
		os.Exit(4)
	}
}
