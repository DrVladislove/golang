package main

import (
	"fmt"
	"os"
)

func main() {
	n := os.Args
	fmt.Printf("Your name is %s and your lastname is %s", n[1], n[2])
}
