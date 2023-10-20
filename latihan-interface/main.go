package main

import "fmt"

type kalkulator interface {
	perkalian() float32
	pembagian() float32
	pengurangan() float32
}

type angka struct {
	a float32
	b float32
}

func inangka(angkake1, angkake2 float32) kalkulator {
	return &angka{a: angkake1, b: angkake2}

}

func (n angka) pengurangan() float32 {
	return n.a - n.b
}
func (n angka) perkalian() float32 {
	return n.a * n.b
}
func (n angka) pembagian() float32 {
	return n.a / n.b
}

func main() {
	masukanangka := inangka(4.0, 2.0)
	fmt.Println(masukanangka.pengurangan())

}
