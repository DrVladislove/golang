package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args
	msg := `hi ` + name[1] + `! 
	how are you?`
	fmt.Println(msg)
}
