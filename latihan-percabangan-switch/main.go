package main

import "fmt"

func main() {
	huruf := "a"

	switch huruf {
	case "a", "i", "u", "e", "o":
		fmt.Printf("Huruf %s adalah huruf vokal", huruf)
	default:
		fmt.Printf("Huruf %s adalah huruf konsonan", huruf)
	}

}
