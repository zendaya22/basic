package main

import (
	"fmt"
)

type objectMenu struct {
	nama string
	harga int
}

func main() {

	var selectMenu int
	var jumlahPesan int
	fmt.Print("==========Selamat Datang=============\n")
	// menu
	var menus = map[int]string{
		1 : "AYAM",
		2 : "BEBEK",
		3 : "KAMBING",
	}
	var harga = map[int]int{
		1 : 20000,
		2 : 15000,
		3 : 40000,
	}
	fmt.Printf("MENU RESTOURANT DIMAS\n1.%v\n2.%v\n3.%v\n", menus[1], menus[2], menus[3])
	// pick 


	fmt.Print("masukan angka 1/2/3 untuk memilih menu : ")
	fmt.Scan(&selectMenu)

	fmt.Print("masukan jumlah pesanan : ")
	fmt.Scan(&jumlahPesan)

	hasilPesanan := objectMenu{
		nama: menus[selectMenu],
		harga: harga[selectMenu],
	}
	fmt.Print("===========PESANAN============\n")
	fmt.Println("NAMA      HARGA     JUMLAH   ")
	result := fmt.Sprintf("%v      %v       %v\n", hasilPesanan.nama, hasilPesanan.harga, jumlahPesan)
	fmt.Println(result)
	fmt.Print("===========TOTAL============\n")
	fmt.Println("total yang harus dibayar : ", hasilPesanan.harga * jumlahPesan)
}
