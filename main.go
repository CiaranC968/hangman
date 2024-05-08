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

	if guess == "_" {
		fmt.Println("Congrats")
	} else {
		print("you lose")
	}
}

func message() string {
	return welcomeMessage
}

func word() string {
	wordList := []string{"ryu", "ken", "gouken", "vega", "chun-li"}
	num := rand.Intn(len(wordList))
	return wordList[num]
}

func repeatString(s string, count int) string {
	return strings.Repeat(s, count)
}
