package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	rndNums = 10
	usage   = `Enter your lucky num
	This program generates %d random nums.
	Try it! Good luck!`
)

func main() {
	rand.Seed(time.Now().Unix())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, rndNums)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Enter number.")
		return
	}

	if guess < 0 {
		fmt.Println("It's negative num, enter positive")
		return
	} else if guess == 0 {
		fmt.Println("It's zero, enter positive num")
	}

	var maxNumber int
	if guess <= 10 {
		maxNumber = 10
	} else {
		maxNumber = guess
	}

	for turn := 1; turn <= rndNums; turn++ {
		n := rand.Intn(maxNumber)

		if n == guess {
			if turn == 1 {
				fmt.Println("You won the first time! You are lucky!")
			} else {
				fmt.Println("YOU WON!")
			}
			return
		}
		fmt.Println(n)
	}

	fmt.Println("YOU LOST. But try again!")
}
