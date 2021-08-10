package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	availableOp = "%*/+-"

	usageNot     = "Usage: [op=" + availableOp + "] [size]"
	missingMsg   = "Size is missing"
	invalidOpMsg = `Invalid operator.
Valid ops one of: ` + availableOp

	invalidOp = -1
)

func main() {
	args := os.Args[1:]

	switch l := len(args); {
	case l == 1:
		fmt.Println(missingMsg)
		fallthrough
	case l < 1:
		fmt.Println(usageNot)
		return
	}

	op := args[0]
	if strings.IndexAny(op, availableOp) == invalidOp {
		fmt.Println(invalidOpMsg)
		return
	}

	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", op)
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {
			var res int

			switch op {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i / j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}

			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
}
