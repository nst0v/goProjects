package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1

	fmt.Print("Guess a number between 0 and 100: ")

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			return
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error parsing input")
			return
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		} else {
			success = true
			fmt.Println("Congratulations! You guessed the number.")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
