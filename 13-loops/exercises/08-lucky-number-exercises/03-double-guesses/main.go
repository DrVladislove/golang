package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 10
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

	guess1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Enter num pls.")
		return
	}

	var guess2 int
	if len(args) == 2 {
		guess2, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Enter num pls.")
			return
		}
	}

	if guess1 <= 0 || guess2 <= 0 {
		fmt.Println("Enter positive num pls.")
		return
	}

	max := guess1
	if guess1 < guess2 {
		max = guess2
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(max) + 1

		if n == guess1 || n == guess2 {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("YOU WON")
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