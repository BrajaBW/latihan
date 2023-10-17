package main

import "fmt"

func main() {
	fruits := map[string]int{

		"Apple":  10,
		"Banana": 15,
		"Grape":  20,
		"Orange": 8,
	}
	fmt.Println("Sebelum ditambah/hapus :\n", fruits)
	fmt.Println("")
	fruits["Manggo"] = 12
	fruits["Strawbery"] = 7
	fmt.Println("Setelah menambahkan Mango dan Strawberry:\n", fruits)
	fmt.Println("")
	delete(fruits, "Orange")
	fmt.Println("Setelah dihapus Orange\n", fruits)
	fmt.Println("")

	key := "Apple"
	val, isExist := fruits["Apple"]

	if isExist {
		fmt.Printf("jumlah %s yang tersedia %d\n", key, val)
	} else {
		fmt.Println(key, "Buah apel tidak ditemukan\n")
	}
	fmt.Println("\n")
	fruts2 := []map[string]string{

		map[string]string{"namabuah": "Banana", "jumlah": "15"},
		map[string]string{"namabuah": "Grape", "jumlah": "20"},
		map[string]string{"namabuah": "Manggo", "jumlah": "12"},
		map[string]string{"namabuah": "Strawbery", "jumlah": "7"},
		map[string]string{"namabuah": "Apple", "jumlah": "10"},
	}

	for _, buah := range fruts2 {

		fmt.Println(buah["namabuah"], ":", buah["jumlah"])
	}

}
