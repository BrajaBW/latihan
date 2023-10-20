package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func (n *Student) SetMyName(ganti string) {
	n.Name = ganti
	fmt.Println("Hello nama saya :", n.Name)

}

func (n Student) CallName() {
	fmt.Println("Nama awal student:", n.Name)
}
func main() {
	name := Student{
		Name:  "Noobee",
		Class: "12A",
	}
	name.CallName()
	name.SetMyName("BRAJA BUANA WIDODO")
}
