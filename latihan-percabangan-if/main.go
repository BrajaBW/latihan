package main

import "fmt"

func main() {
	num := 4

	if num%2 != 0 {

		fmt.Printf("%d bilangan ganjil", num)
	} else {
		fmt.Printf("%d bilangan genap\n", num)
	}
	num2 := 5

	if num2%2 != 0 {

		fmt.Printf("%d bilangan ganjil", num2)
	} else {
		fmt.Printf("%d bilangan genap", num2)
	}
}
