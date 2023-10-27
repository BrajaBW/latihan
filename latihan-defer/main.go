package main

import "fmt"

func main() {

	num1 := 5
	num2 := 5
	jumlah := num1 + num2

	fmt.Println("masukan angka Pertama:",num1)
	fmt.Println("masukan angka Kedua:",num2)
	fmt.Println("Hasil penjumlahan:",jumlah)
func(){
	defer cetak("Program Selesai")
}()
}

func cetak(text string)  {
	fmt.Println(text)
	
}
