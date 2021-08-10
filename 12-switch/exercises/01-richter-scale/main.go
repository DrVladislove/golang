package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Give me the magnitude of the earthquake.")
		return
	}

	magnitude, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("I couldn't get that&& sorry.")
		return
	}

	switch m := magnitude; {
	case m < 2:
		fmt.Printf("%.2f is micro \n", m)
	case m >= 2 && m < 3:
		fmt.Printf("%.2f is very minor \n", m)
	case m >= 3 && m < 4:
		fmt.Printf("%.2f is minor \n", m)
	case m >= 4 && m < 5:
		fmt.Printf("%.2f is light \n", m)
	case m >= 5 && m < 6:
		fmt.Printf("%.2f is moderate \n", m)
	case m >= 6 && m < 7:
		fmt.Printf("%.2f is strong \n", m)
	case m >= 7 && m < 8:
		fmt.Printf("%.2f is major \n", m)
	case m >= 8 && m < 10:
		fmt.Printf("%.2f is great \n", m)
	default:
		fmt.Printf("%.2f is massive", m)
	}
}
