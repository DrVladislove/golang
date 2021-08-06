package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days := 28

	month := os.Args[1]

	if mon := strings.ToLower(month); mon == "june" ||
		mon == "april" ||
		mon == "september" ||
		mon == "november" {
		days = 30
	} else if mon == "january" ||
		mon == "march" ||
		mon == "may" ||
		mon == "july" ||
		mon == "august" ||
		mon == "october" ||
		mon == "december" {
		days = 31
	} else if mon == "february" {
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
