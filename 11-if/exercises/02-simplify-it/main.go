package main

import "fmt"

func main() {
	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Printf("It's a big sphere.")
	} else {
		fmt.Printf("I don't know.")
	}
}
