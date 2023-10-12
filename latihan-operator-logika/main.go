package main

import "fmt"

func main() {
	gender := "famele"
	umur := 20
	if gender == "famele" && umur >= 21 {
		fmt.Println("Pelamar dapat dipekerjakan")
	} else {
		fmt.Println("Pelamar tidak dapat dipekerjakan")
	}
}
