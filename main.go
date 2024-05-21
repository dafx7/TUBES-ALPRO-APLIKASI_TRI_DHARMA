package main

import "fmt"

const NMAX int = 30

type PPM struct {
	ketua, prodi, judul, sumber_dana, luaran, tahun_kegiatan string
	anggota                                                  [4]string
}

type arrPPM [NMAX]PPM

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
		if pilihan == 1 {
			tambah_data()
		} else if pilihan == 2 {
			edit_data()
		} else if pilihan == 3 {
			hapus_data()
		} else if pilihan == 4 {
			tampilkan_data()
		} else if pilihan == 5 {
			urutkan_data()
		}
		tulisan_menu()
		fmt.Scan(&pilihan)
	}
	fmt.Print("Terimakasih telah menggunakan aplikasi Tri Dharma Perguruan Tinggi.")
}

func tambah_data() {

}

func edit_data() {

}

func hapus_data() {

}

func tampilkan_data() {

}

func urutkan_data() {

}
