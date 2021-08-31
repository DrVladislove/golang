package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		fmt.Println("wrong numbers")
		return
	}

	max, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		fmt.Println("wrong numbers")
		return
	}

	if min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		if i%2 == 0 {
			sum += i

			fmt.Print(i)
			if i < max {
				fmt.Print(" + ")
			}
		}

	}
	fmt.Printf(" = %d\n", sum)
}