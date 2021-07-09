package main

import (
	"fmt"
	"os"
)

func main() {
	//Main task
	fmt.Println(os.Args[1])

	//Bonus task
	fmt.Println("Hello,", os.Args[1])
	fmt.Println("How are you?")
}
