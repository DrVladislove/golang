package main

import (
	"fmt"
	"os"
)

const (
	usage           = "Usage: [username] [password]"
	errAccessDenied = "Access denied for %q.\n"
	errPwd          = "Invalid password for %q.\n"
	accessOK        = "Access granted to %q.\n"
	user            = "jack"
	pass            = "1888"
	user2           = "inanc"
	password2       = "1879"
)

type CheckAssessResult int

const (
	Allowed CheckAssessResult = 1 + iota
	InvalidPassword
	AccessDenied
)

func checkAccess(username, password string, desiredUsername, desiredPassword string) CheckAssessResult {
	if username == desiredUsername && password == desiredPassword {
		return Allowed
	}
	if username == desiredUsername && password != desiredPassword {
		return InvalidPassword
	}
	return AccessDenied
}

func printResult(result CheckAssessResult, username string) {

	if result == Allowed {
		fmt.Println(accessOK, username)
	}
	if result == InvalidPassword {
		fmt.Println(errPwd, username)
	}
	fmt.Println(errAccessDenied, username)
}

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	result := checkAccess(u, p, user, pass)
	if result != AccessDenied {
		printResult(result, u)
		return
	}

	result = checkAccess(u, p, user2, password2)
	printResult(result, u)
}
