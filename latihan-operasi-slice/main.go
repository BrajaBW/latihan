package main

import "fmt"

func main() {
	myslice := []int{10, 20, 30, 40, 50}
	fmt.Println("====[mySlice]====")
	fmt.Println("Data :", myslice)
	fmt.Println("Len :", len(myslice))
	fmt.Println("Cap :", cap(myslice))
	fmt.Println("")
	slice1 := myslice[0:3]
	fmt.Println("====[Slice1]====")
	fmt.Println("Data :", slice1)
	fmt.Println("Len :", len(slice1))
	fmt.Println("Cap :", cap(slice1))
	fmt.Println("")
	slice2 := myslice[3:5]
	fmt.Println("====[Slice2]====")
	fmt.Println("Data :", slice2)
	fmt.Println("Len :", len(slice2))
	fmt.Println("Cap :", cap(slice2))
	fmt.Println("")
	dataApend1 := append(slice1, 60)
	fmt.Println("====[Setelah Append ke Slice1]====")
	fmt.Println("data Slice 1:", dataApend1)
	fmt.Println("Len Slice 1::", len(dataApend1))
	fmt.Println("Cap Slice 1::", cap(dataApend1))
	fmt.Println("Data Slice 2:", slice2)
	fmt.Println("Len Slice 2:", len(slice2))
	fmt.Println("Cap Slice 2:", cap(slice2))
	fmt.Println("")
	dataApend2 := append(slice2, 70)
	fmt.Println("====[Setelah Append ke Slice2]====")
	fmt.Println("data Slice 1:", dataApend1)
	fmt.Println("Len Slice 1:", len(dataApend1))
	fmt.Println("Cap Slice 1:", cap(dataApend1))
	fmt.Println("Data Slice 2:", dataApend2)
	fmt.Println("Len Slice 2:", len(dataApend2))
	fmt.Println("Cap Slice 2 :", cap(dataApend2))

}
