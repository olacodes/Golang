package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("\n" + "===============Guessing Game===================" + "\n")

	fmt.Println("I have chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	computerGuess := rand.Intn(101)

	x := 0
	for x < 10 {
		x++
		fmt.Print("Make a guess: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		userGuess, _ := strconv.Atoi(input)
		if userGuess > computerGuess {
			fmt.Println("Oops your guess too HIGH")
			msg(x)
		} else if userGuess < computerGuess {
			fmt.Println(("Oops You guess too LOW"))
			msg(x)
		} else {
			msg := fmt.Sprintf("You guessed it. After %d times", x)
			fmt.Println(msg + "\n")
			break
		}
	}
	fmt.Println("You loose. You have exhausted all guesses")
}

func msg(x int) {
	msg := fmt.Sprintf("You've guessed %d times and your remaining guess is %d", x, 10-x)
	fmt.Println(msg + "\n")
}