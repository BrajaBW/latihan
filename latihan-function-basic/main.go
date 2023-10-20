package main

import (
	"fmt"
)

func kata() (a string) {
	a = "Mobil"
	return a
}

func cetak(b string) string {
	return b
}

func main() {

	var car = map[string]string{}
	car["name"] = "BWM"
	car["color"] = "Black"

	for _, mobil := range car {
		cetak(mobil)
	}

	fmt.Println(kata(), cetak(car["name"]), cetak(car["color"]))

}
