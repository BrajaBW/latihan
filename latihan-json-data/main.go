package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Nama  string
	Class string
}

func saveToFile(users []User) {
	userjeson, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("bentuk Json :", string(userjeson))
}

func loadFromFile ()

func main() {

	users:=[]User{
	User{Nama : "Noob", Class: "A"},
	User{Nama : "Java", Class: "B"},
	User{Nama : "Docker", Class: "C"},

	
	}
err:= 
	fmt.Println(saveToFile(users{}))
	saveToFile()


}
