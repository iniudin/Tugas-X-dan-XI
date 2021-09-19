package main

import (
	"errors"
	"fmt"
	"unicode"
)

func getLetter(word string, length int) string {
	result := ""
	if length == 0 {
		return result
	} else {
		if unicode.IsLetter(rune(word[length-1])) {
			result += string(word[length-1])
			return result + getLetter(word, length-1)
		}
		return getLetter(word, length-1)
	}

}

func Solver(word string) (string, error) {
	result := getLetter(word, len(word))
	if len(result) == 0 {
		fmt.Println(result)
		return result, errors.New("Harus terdapat string")
	}

	return result, nil
}

func main() {
	var inputUser string

	fmt.Println("Input problem: ")
	fmt.Scan(&inputUser)

	result, err := Solver(inputUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}
