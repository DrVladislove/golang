package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max := os.Args

	if len(max) != 2 {
		fmt.Println("Give me the size of the table")
		return
	}

	s, err := strconv.Atoi(max[1])
	if err != nil || s <= 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= s; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= s; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= s; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
