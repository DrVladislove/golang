package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Give a year number")
		return
	}

	y, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a valid year\n", os.Args[1])
		return
	}

	if y%100 != 0 && y%4 == 0 {
		fmt.Printf("%d is a lip year", y)
	} else {
		fmt.Printf("%d is not a leap year", y)
	}
}
