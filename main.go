package main

import "fmt"

const NMAX int = 30

type PPM struct {
	jenis, ketua, prodi, judul, sumber_dana, luaran string
	tahun_kegiatan, jumAnggota                      int
	anggota                                         [4]string
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
	var ArrayPPM arrPPM
	var nPPM int

	tulisan_menu()
	fmt.Scan(&pilihan)
	for pilihan != 6 {
		if pilihan == 1 {
			tambah_data(&ArrayPPM, &nPPM)
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

func tambah_data(A *arrPPM, n *int) {
	var nAnggota int
	// Input Jenis PPM
	fmt.Print("Input jenis PPM: ")
	fmt.Scan(&A[*n].jenis)

	fmt.Print("Input nama ketua: ")
	fmt.Scan(&A[*n].ketua)
	fmt.Print("Input jumlah anggota: ")
	fmt.Scan(&nAnggota)

	// Jika inputan pengguna itu lebih dari 4, maka pengguna akan diminta untuk input ulang.
	for nAnggota > 4 {
		fmt.Print("WARNING!!")
		fmt.Println("Jumlah anggota melebihi batas (Max 4)")
		fmt.Print("Input jumlah anggota: ")
		fmt.Scan(&nAnggota)
	}
	// Input nama anggota sebanyak nAnggota.
	for i := 0; i < nAnggota; i++ {
		fmt.Printf("Input anggota ke-%d: ", i+1)
		fmt.Scan(&A[*n].anggota[i])
	}
	A[*n].jumAnggota = nAnggota

	fmt.Print("Input prodi: ")
	fmt.Scan(&A[*n].prodi)
	fmt.Print("Input judul: ")
	fmt.Scan(&A[*n].judul)
	fmt.Print("Input sumber dana: ")
	fmt.Scan(&A[*n].sumber_dana)
	fmt.Print("Input luaran PPM: ")
	fmt.Scan(&A[*n].luaran)
	fmt.Print("Input tahun kegiatan: ")
	fmt.Scan(&A[*n].tahun_kegiatan)
}

func edit_data() {

}

func hapus_data() {

}

func tampilkan_data() {

}

func urutkan_data() {

}
