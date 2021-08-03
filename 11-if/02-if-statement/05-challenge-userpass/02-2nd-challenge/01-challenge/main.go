package main

import (
	"fmt"
	"os"
)

const (
	usage     = "Usage: [username] [password]"
	errUser   = "Access denied for %q.\n"
	errPwd    = "Invalid password for %q.\n"
	accessOK  = "Access granted to %q.\n"
	user      = "jack"
	pass      = "1888"
	user2     = "inanc"
	password2 = "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if p != pass && p != password2 {
		fmt.Printf(errPwd, u)
	} else if u == user && p == pass {
		fmt.Printf(accessOK, u)
	} else if u == user2 && p == password2 {
		fmt.Printf(accessOK, u)
	}
}
