package main

import (
	"fmt"
)

// panic
func validation(numb1, numb2 int) {
	defer cegahError()
	if numb1 == 0 || numb2 == 0 {
		fmt.Println("Masukkan angka pertama:", numb1)
		fmt.Println("Masukkan angka kedua:", numb2)
		panic("tidak bisa membagi data dengan 0")
	} else {
		hasil := numb1 / numb2
		fmt.Println("Masukkan angka pertama:", numb1)
		fmt.Println("Masukkan angka kedua:", numb2)
		fmt.Println("Hasil pembagian:", hasil)
	}
}

func cegahError() {

	err := recover()
	if err != nil {

		fmt.Println("RECOVER", err)
	}

}

func main() {

	validation(0, 3)
}
