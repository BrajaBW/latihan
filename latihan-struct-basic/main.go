package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func main() {
	var Book1 Book
	Book1.Title = "Pemrograman Go"
	Book1.Author = "John Doe"

	fmt.Printf("Informasi tentang Book 1:\n Judul:%s\n Penulis:%s\n", Book1.Author, Book1.Title)
	fmt.Println("")
	Book2 := Book{
		Title:  "Algoritma dan Struktur Data",
		Author: "Jane Smith",
	}
	fmt.Printf("Informasi tentang Book 2:\n Judul:%s\n Penulis:%s\n", Book2.Title, Book2.Author)
	fmt.Println("")

	Book3 := Book{"Machine Learning for Beginners", "David Johnson"}
	fmt.Printf("Informasi tentang Book 3:\n Judul:%s\n Penulis:%s", Book3.Author, Book3.Title)

}
