package main

import (
	"fmt"
)

func main() {
	// // Latihan 1

	for counter := 5; counter >= 1; counter-- {
		for counters := 1; counters <= counter; counters++ {

			fmt.Print("* ")

		}
		fmt.Println()

	}
	//lATIHAN 2
	for counter := 1; counter <= 5; counter++ {
		for counters := 1; counters <= counter; counters++ {

			fmt.Print("* ")

		}
		fmt.Println()
	}
	// Latihan 3
	for i := 1; i <= 50; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("NooBee")
		} else if i%3 == 0 {
			fmt.Println("Noo")
		} else if i%5 == 0 {
			fmt.Println("Bee")
		} else {
			fmt.Println(i)
		}
	}
}
