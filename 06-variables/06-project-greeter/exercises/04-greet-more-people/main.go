package main

import (
	"fmt"
	"os"
)

func main() {
	count := len(os.Args) - 1
	fmt.Println("There are", count, "people!")
	fmt.Println("Hello great", os.Args[1], "!")
	fmt.Println("Hello great", os.Args[2], "!")
	fmt.Println("Hello great", os.Args[3], "!")
	fmt.Println("Nice to meet you all.")
}
