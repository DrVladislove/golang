package main

import "fmt"

func main() {
	json := `{
		"Items": [{
		"Item": {
		"name": "Teddy Bear"
		}
		}]
	}`
	fmt.Println(json)
}
