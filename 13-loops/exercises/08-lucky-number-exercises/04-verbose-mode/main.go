package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 7
	usage    = `Welcome to the My Game! 
The program will generate %d random nums.
You need guess one of generated nums, but remember that the greater your number, harder to win.
Try it right now!
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var show string
	if args[0] == "-v" {
		show = "-v"
	}

	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("Enter num pls.")
		return
	}

	if guess <= 0 {
		fmt.Println("Enter positive num pls.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess) + 1
		if show == "-v" {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			switch rand.Intn(3) {
			case 0:

			case 1:
				fmt.Println("YOU'RE AWESOME")
			case 2:
				fmt.Println("GOOD JOB!")
			}
			return
		}

	}

	try := "%s Try again?\n"

	switch rand.Intn(2) {
	case 0:
		fmt.Printf(try, "LOSER!")
	case 1:
		fmt.Printf(try, "YOU LOST")
	}

}
