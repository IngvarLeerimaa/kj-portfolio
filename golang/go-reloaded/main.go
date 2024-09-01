package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {

	fileName := os.Args[1]
	sampletxt, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	// to print it out in terminal
	//fmt.Println(string(sampletxt))

	result := work(sampletxt)

	os.WriteFile("result.txt", result, 0644)
	// to print it out in terminal
	//fmt.Println(string(result))

}

func work(sampletxt []byte) []byte {

	text := string(sampletxt)

	//splits the given string around each instance of one or more consecutive white space characters

	slice := strings.Fields(text)

	//check for singles in txt and convert previous string

	for i := 0; i < len(slice); i++ {
		if slice[i] == "(hex)" {
			num, _ := strconv.ParseInt(slice[i-1], 16, 64)
			convNum := int(num)
			slice[i-1] = strconv.Itoa(convNum)
			slice = append(slice[:i], slice[i+1:]...)

		}
		if slice[i] == "(bin)" {
			num, _ := strconv.ParseInt(slice[i-1], 2, 64)
			convNum := int(num)
			slice[i-1] = strconv.Itoa(convNum)
			slice = append(slice[:i], slice[i+1:]...)

		}
		if slice[i] == "(up)" {
			word := strings.ToUpper(slice[i-1])
			slice[i-1] = word
			slice = append(slice[:i], slice[i+1:]...)
		}
		if slice[i] == "(low)" {
			word := strings.ToLower(slice[i-1])
			slice[i-1] = word
			slice = append(slice[:i], slice[i+1:]...)
		}
		if slice[i] == "(cap)" {
			word := strings.Title(slice[i-1]) //its yellow because we have to use standard packages and cant use Caser
			slice[i-1] = word
			slice = append(slice[:i], slice[i+1:]...)
		}

		// check for txt with more than 1 strings to convert

		if slice[i] == "(low," {
			word2 := strings.TrimRight(slice[i+1], ")")
			num2, _ := strconv.Atoi(word2)
			if num2 > 0 {
				for j := num2; j > 0; j-- {
					word := strings.ToLower(slice[i-j])
					slice[i-j] = word
				}
				slice = append(slice[:i], slice[i+2:]...)
			}
		}
		if slice[i] == "(up," {
			word2 := strings.TrimRight(slice[i+1], ")")
			num2, _ := strconv.Atoi(word2)
			if num2 > 0 {
				for j := num2; j > 0; j-- {
					word := strings.ToUpper(slice[i-j])
					slice[i-j] = word
				}
				slice = append(slice[:i], slice[i+2:]...)
			}
		}
		if slice[i] == "(cap," {
			word2 := strings.TrimRight((slice[i+1]), ")")
			num2, _ := strconv.Atoi(word2)
			if num2 > 0 {
				for j := num2; j > 0; j-- {
					word := strings.Title(slice[i-j])
					slice[i-j] = word
				}
				slice = append(slice[:i], slice[i+2:]...)
			}
		}

		//spelling

		if slice[i] == "A" || slice[i] == "a" {
			nextWord := slice[i+1]
			vowels := "aeiouh"
			for k := 0; k < len(vowels); k++ {
				if rune(nextWord[0]) == rune(vowels[k]) {
					slice[i] = slice[i] + "n"
				}
			}
		}

	}

	//punctuation
	convWord := strings.Join(slice, " ")
	runes := []rune(convWord)

	for j := 0; j < len(convWord); j++ {

		if runes[j] == '.' && runes[j-1] == ' ' || runes[j] == ',' && runes[j-1] == ' ' || runes[j] == '!' && runes[j-1] == ' ' || runes[j] == '?' && runes[j-1] == ' ' || runes[j] == ':' && runes[j-1] == ' ' || runes[j] == ';' && runes[j-1] == ' ' {
			runes[j], runes[j-1] = runes[j-1], runes[j] // for switcing places between punctuation and space
		} else if runes[j] == 32 && runes[j-1] == 32 { // for deleting double spacing
			runes[j-1] = 127
		}
	}

	quotes(runes)

	convWord = string(runes)

	// used acii 127 ((del) in ln 129, 157 and 159) an it left a unicode char to result.txt when viewed in vscode.
	convWord = strings.Replace(convWord, "", "", -1)
	return []byte(convWord)
}
func quotes(runes []rune) []rune {
	runesx := runes
	k := 0
	m := 0

	for i := 0; i < len(runes); i++ {

		if runes[i] == 39 {
			k++
		} else if k < 1 {
			m++
		}
		for j := 0; j < len(runes); j++ {

			if k == 2 && runes[j] == ' ' {

				runesx[i-1] = 127
			} else if k == 1 && (runes[j] == 32) {
				runesx[m+1] = 127

			}
		}

	}
	return runesx
}
