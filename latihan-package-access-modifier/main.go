package main

import(
	"fmt"
	"latihan-package-access-modifier/calculator"

)




func main() {
	num1:=1
	num2:=2

	penjumlahan := calculator.Add(num1,num2)
	// Perkalian := calculator.perkalian(num1,num2) <func harus HUrup Besar Ketika di panggil beda package>
	fmt.Printf("hasil Penjumlahan:%d\n",penjumlahan)
	// fmt.Printf("hasil Penjumlahan:%d",Perkalian)
	
}
