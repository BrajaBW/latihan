package main

import (
	"fmt"
	"net/http"
)

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("about called.....")
	tentang := "ini Adalah halaman About"
	w.Write([]byte(tentang))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index called.....")
	halaman := "ini Adalah halaman index"
	w.Write([]byte(halaman))
}e

func handleRoute() {
	http.HandleFunc("/about", about)
	http.HandleFunc("/index", index)
}

func starServer(port string) {
	handleRoute()
	fmt.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

func main() {

	port := ":4345"
	starServer(port)

}
