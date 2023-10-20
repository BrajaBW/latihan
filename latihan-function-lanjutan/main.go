package main

import "fmt"

func main() {
	nilaisiswa := []float32{85.5, 92.0, 78.5, 90.0, 88.5}
	total := tot(nilaisiswa...)
	fmt.Println("Jumlah nilai siswa dalam kelas:", len(nilaisiswa))
	pembagian := len(nilaisiswa)
	fmt.Printf("Nilai rata-rata siswa dalam kelas:%.2f", total/float32(pembagian))

}

func tot(r ...float32) (total float32) {
	for _, a := range r {
		total += a

	}
	return total

}
