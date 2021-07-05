package main

import (
	"fmt"
	"path"
)

func main() {
	a, _ := path.Split("secret/file.txt")
	fmt.Println(a)
}
