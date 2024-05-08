package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const welcomeMessage = "Welcome! Let's play Hangman!"

func main() {
	fmt.Println(message())
	fmt.Println("________")
	fmt.Println("|")
	fmt.Println("|")
	fmt.Print("|")
	fmt.Println(repeatString("_", len(word())))

	fmt.Print("Enter Your guess: ")

	var guess string

	// Taking input from user
	fmt.Scan(&guess)

	if containsLetter(word(), guess) {
		fmt.Printf("The word contains the letter '%s'.\n", guess)
	} else {
		fmt.Printf("The word does not contain the letter '%s'.\n", guess)
	}
}

func message() string {
	return welcomeMessage
}

func containsLetter(word, letter string) bool {
	return strings.Contains(word, letter)
}

func word() string {
	wordList := []string{"ryu", "ken", "gouken", "vega", "chun-li"}
	num := rand.Intn(len(wordList))
	return wordList[num]
}

func repeatString(s string, count int) string {
	return strings.Repeat(s, count)
}
