package main

import "fmt"

func main() {
	fmt.Println("Selamat datang di aplikasi Tri Dharma Perguruan Tinggi.")
	menu_utama()
}

func tulisan_menu() {
	fmt.Println("1. Menambahkan Data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Tampilkan Data")
	fmt.Println("5. Urutkan Data")
	fmt.Println("6. Keluar")

	fmt.Print("Input pilihan: ")
}

func menu_utama() {
	var pilihan int

	tulisan_menu()
	fmt.Scan(&pilihan)
	for pilihan != 6 {
		tulisan_menu()
		fmt.Scan(&pilihan)
	}
	fmt.Print("Terimakasih telah menggunakan aplikasi Tri Dharma Perguruan Tinggi.")
}
