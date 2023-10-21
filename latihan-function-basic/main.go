package main

import (
	"fmt"
)

// func kata() (a string) {
// 	a = "Mobil"
// 	return a
// }

// func cetak(b string) string {
// 	return b
// }

func combineString(str1 string, str2 string) string {
	return "Mobil " + str1 + " Berwarna " + str2
}

func cetak(str3 string) {
	fmt.Println(str3)

}

func main() {

	var car = map[string]string{}
	car["name"] = "BWM"
	car["color"] = "Black"

	cetaklangsung := combineString(car["name"], car["color"])
	cetak(cetaklangsung)

	// for _, mobil := range car {
	// 	cetak(mobil)
	// }

	// fmt.Println(kata(), cetak(car["name"]), cetak(car["color"]))

}
