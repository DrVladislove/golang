package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	username, password := args[1], args[2]

	if username == "jack" && password == "1888" {
		fmt.Printf("Access granted to %q. \n", username)
	} else if username != "jack" && password != "1888" {
		fmt.Printf("Access denied for %q. \n", username)
	} else if username == "jack" && password != "1888" {
		fmt.Printf("Invalid password for %q. \n", username)
	}
}
