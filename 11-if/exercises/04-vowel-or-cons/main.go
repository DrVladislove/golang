package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
		return
	}

	let := args[1]
	if let == "w" || let == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", let)
	} else if let == "a" || let == "e" || let == "i" || let == "o" || let == "u" {
		fmt.Printf("%q is a vowel.\n", let)
	} else {
		fmt.Printf("%q is a consonant.\n", let)
	}
}
