package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const welcomeMessage = "Welcome! Let's play Hangman!"

func main() {
	fmt.Println(message())
	fmt.Println("________")
	fmt.Println("|")
	fmt.Println("|")
	fmt.Print("|")
	fmt.Println(repeatString("_", len(word())))

}

func message() string {
	return welcomeMessage
}

func word() string {
	wordList := []string{"ryu", "ken", "gouken", "vega", "chun-li"}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(wordList))
	return wordList[num]
}

func repeatString(s string, count int) string {
	return strings.Repeat(s, count)
}
