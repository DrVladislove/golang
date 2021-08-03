package main

import "fmt"

func main() {
	age := 10

	if age > 60 {
		fmt.Printf("Getting older")
	} else if age > 30 && age < 60 {
		fmt.Printf("Getting wiser")
	} else if age > 20 && age < 30 {
		fmt.Printf("Adulthood")
	} else if age > 10 && age < 20 {
		fmt.Printf("Young blood")
	} else {
		fmt.Printf("Booting up")
	}
}
