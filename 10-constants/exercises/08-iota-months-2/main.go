package main

import (
	"fmt"
)

func main() {
	const (
		_   = iota
		Jan = iota
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)
}
