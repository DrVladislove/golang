package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	msg := os.Args[1]
	len := utf8.RuneCountInString(msg)

	s := msg + strings.Repeat("!", len)

	fmt.Println(s)
}
