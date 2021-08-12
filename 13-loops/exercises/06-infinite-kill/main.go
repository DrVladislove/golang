package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		var char string

		switch rand.Intn(4) {
		case 0:
			char = "|"
		case 1:
			char = "-"
		case 2:
			char = "|"
		case 3:
			char = "\\"
		}
		fmt.Printf("\r%s Please Wait. Processing....", char)
		time.Sleep(time.Millisecond * 250)
	}
}
