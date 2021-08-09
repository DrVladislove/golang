package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Tell me the magnitude of the earthquake in human terms.")
		return
	}
	m := os.Args[1]
	switch m {
	case "micro":
		fmt.Printf("%q's  richter scale is less than 2.0", m)
	case "very minor":
		fmt.Printf("%q's  richter scale is 2 - 2.9", m)
	case "minor":
		fmt.Printf("%q's  richter scale is 3 - 3.9", m)
	case "light":
		fmt.Printf("%q's  richter scale is 4 - 4.9", m)
	case "moderate":
		fmt.Printf("%q's  richter scale is 5 - 5.9", m)
	case "strong":
		fmt.Printf("%q's  richter scale is 6 - 6.9", m)
	case "major":
		fmt.Printf("%q's  richter scale is 7 - 7.9", m)
	case "great":
		fmt.Printf("%q's  richter scale is 8 - 8.9", m)
	case "massive":
		fmt.Printf("%q's  richter scale is 10+", m)
	default:
		fmt.Printf("%q's richter scale is unknown", m)
	}
}
