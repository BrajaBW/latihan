package main

import "fmt"

func printFibonacciSeries(int){
	a := 0
	b := 1
	c := b
	fmt.Printf("%d %d", a, b)
	for true{
	   c = b
	   b = b+ 2
	   if b >= 100{
		  fmt.Println()
		  break
	   }
	   a = c
	   fmt.Printf(" %d", b)
	}
 }
 
 func main(){
	printFibonacciSeries(10)
 }