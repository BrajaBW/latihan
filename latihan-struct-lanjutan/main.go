package main

import "fmt"

type Parents struct {
	Nama string
	Umur int
}

type Siswa struct {
	OrangTua Parents
	Nama     string
	Umur     int
	Kelas    string
}

type DaftarSiswa struct {
	Nama         string
	Umur         int
	Kelas        string
	NamaOrangTua string
	UmurOrangTua int
}

func main() {
	orangtua1 := Parents{
		Nama: "Asep",
		Umur: 55,
	}
	siswa1 := Siswa{
		OrangTua: orangtua1,
		Nama:     "Ali",
		Umur:     15,
		Kelas:    "9A",
	}
	fmt.Println("Informasi Siswa 1:")
	fmt.Printf("Nama: %s umur: %d Kelas:%s\n Orang Tua:%s Umur:%d\n", siswa1.Nama, siswa1.Umur, siswa1.Kelas, siswa1.OrangTua.Nama, orangtua1.Umur)
	fmt.Println("")

	orangtua2 := Parents{
		Nama: "Udin",
		Umur: 60,
	}
	siswa2 := Siswa{
		OrangTua: orangtua2,
		Nama:     "David",
		Umur:     14,
		Kelas:    "8B",
	}
	fmt.Println("Informasi Siswa 2:")
	fmt.Printf("Nama: %s umur: %d Kelas:%s\n Orang Tua:%s Umur:%d\n", siswa2.Nama, siswa2.Umur, siswa2.Kelas, siswa2.OrangTua.Nama, orangtua2.Umur)
	fmt.Println("")

	orangtua3 := Parents{
		Nama: "Karla",
		Umur: 66,
	}
	siswa3 := Siswa{
		OrangTua: orangtua3,
		Nama:     "Fina",
		Umur:     16,
		Kelas:    "6B",
	}
	fmt.Println("Informasi Siswa 3:")
	fmt.Printf("Nama: %s umur: %d Kelas:%s\n Orang Tua:%s Umur:%d\n", siswa3.Nama, siswa3.Umur, siswa3.Kelas, siswa3.OrangTua.Nama, orangtua2.Umur)
	fmt.Println("")
	var daftarsiswa = []DaftarSiswa{
		{Nama: "Eva", Umur: 12, Kelas: "6B", NamaOrangTua: "Rudi", UmurOrangTua: 40},
		{Nama: "Faisal", Umur: 11, Kelas: "5A", NamaOrangTua: "Dewi", UmurOrangTua: 38},
		{Nama: "Grace", Umur: 10, Kelas: "4C", NamaOrangTua: "Hendro", UmurOrangTua: 42},
	}
	fmt.Println("Daftar nama :")
	for _, value := range daftarsiswa {
		fmt.Printf("Nama :%s, Umur :%d,kelas :%s,\n Orang Tua :%s, Umur Orang Tua :%d\n\n", value.Nama, value.Umur, value.Kelas, value.NamaOrangTua, value.UmurOrangTua)
	}

}
