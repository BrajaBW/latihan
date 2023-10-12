package main

import "fmt"

func main() {
	kontak := make(map[string]int)

	kontak["Alice"] = 1234567890
	kontak["Bob"] = 9876543210
	fmt.Println(kontak)
	fmt.Println("Nomor telepon Alice:", kontak["Alice"])
	fmt.Println("")
	kontak["Charlie"] = 5555555555
	fmt.Println("Setelah menambah kontak Charlie:", kontak)
	fmt.Println("")
	delete(kontak, "Bob")
	fmt.Println("Setelah menambah kontak Bob Dihapus:", kontak)
	fmt.Println("")
	fmt.Println("Data Kontak:")
	for key, value := range kontak {

		fmt.Println(key, value)
	}

}
