package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Requires age")
		return
	}

	// age := args[1]

	age, err := strconv.Atoi(os.Args[1])

	if err != nil || age < 0 {
		fmt.Printf(`Wrong age: %q`+"\n", os.Args[1])
	}

	if age > 17 {
		fmt.Println("R-Rated")
	}
	if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	}
	if age > 0 && age < 13 {
		fmt.Println("PG-Rated")
	}
}
