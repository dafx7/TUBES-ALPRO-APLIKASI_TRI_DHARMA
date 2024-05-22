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
			fmt.Println()
			tambah_data(&ArrayPPM, &nPPM)
		} else if pilihan == 2 {
			fmt.Println()
			edit_data(&ArrayPPM, &nPPM)
		} else if pilihan == 3 {
			fmt.Println()
			hapus_data()
		} else if pilihan == 4 {
			fmt.Println()
			tampilkan_data()
		} else if pilihan == 5 {
			fmt.Println()
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
	fmt.Print("Berapa jumlah anggota: ")
	fmt.Scan(&nAnggota)

	// Jika inputan pengguna itu lebih dari 4, maka pengguna akan diminta untuk input ulang.
	for nAnggota > 4 {
		fmt.Println()
		fmt.Println("WARNING!!")
		fmt.Println("Jumlah anggota melebihi batas (Max 4)")
		fmt.Println()
		fmt.Println("Input nama ketua:", A[*n].ketua)
		fmt.Print("Input jumlah anggota: ")
		fmt.Scan(&nAnggota)
	}
	// Input nama anggota sebanyak nAnggota.
	for i := 0; i < nAnggota; i++ {
		fmt.Printf("Anggota ke-%d: ", i+1)
		fmt.Scan(&A[*n].anggota[i])
	}
	A[*n].jumAnggota = nAnggota

	fmt.Print("Fakultas: ")
	fmt.Scan(&A[*n].prodi)
	fmt.Print("Judul: ")
	fmt.Scan(&A[*n].judul)
	fmt.Print("Sumber dana: ")
	fmt.Scan(&A[*n].sumber_dana)
	fmt.Print("Luaran PPM: ")
	fmt.Scan(&A[*n].luaran)
	fmt.Print("Tahun kegiatan: ")
	fmt.Scan(&A[*n].tahun_kegiatan)
	*n++
}

func edit_data(A *arrPPM, n *int) {
	var pengubah, jenis, judul string
	var idx, idx_anggota, peng_tahun, pilihan int
	fmt.Println("Input filter untuk data yang ingin diubah")
	fmt.Print("Jenis data: ")
	fmt.Scan(&jenis)
	fmt.Print("Judul data: ")
	fmt.Scan(&judul)

	idx = sequential_search(*A, *n, jenis, judul)
	if idx == -1 {
		fmt.Print("Data yang ingin diedit tidak di temukan.")
		return
	}
	fmt.Println("Tentukan jenis data yang mau diubah: ")

	fmt.Println("1. Ketua")
	fmt.Println("2. Anggota")
	fmt.Println("3. Fakultas")
	fmt.Println("4. Judul")
	fmt.Println("5. Sumber dana")
	fmt.Println("6. Luaran PPM")
	fmt.Println("7. Tahun kegiatan")

	fmt.Print("Input pilihan: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		fmt.Print("Masukan nama ketua yang baru: ")
		fmt.Scan(&pengubah)
		A[idx].ketua = pengubah
	} else if pilihan == 2 {
		fmt.Print("Anggota ke berapa yang ingin diedit: ")
		fmt.Scan(&idx_anggota)
		fmt.Print("Input nama baru: ")
		fmt.Scan(&pengubah)
		A[idx].anggota[idx_anggota-1] = pengubah
	} else if pilihan == 3 {
		fmt.Print("Input nama Fakultas yang baru: ")
		fmt.Scan(&pengubah)
		A[idx].prodi = pengubah
	} else if pilihan == 4 {
		fmt.Print("Input judul yang baru: ")
		fmt.Scan(&pengubah)
		A[idx].judul = pengubah
	} else if pilihan == 5 {
		fmt.Print("Input Sumber dana yang baru: ")
		fmt.Scan(&pengubah)
		A[idx].sumber_dana = pengubah
	} else if pilihan == 6 {
		fmt.Print("Input Luaran yang baru: ")
		fmt.Scan(&pengubah)
		A[idx].luaran = pengubah
	} else if pilihan == 7 {
		fmt.Print("Input tahun kegiatan yang baru: ")
		fmt.Scan(&peng_tahun)
		A[idx].tahun_kegiatan = peng_tahun
	}
}

func sequential_search(A arrPPM, n int, jenis, judul string) int {
	var idx int = -1
	for i := 0; i < n; i++ {
		if A[i].jenis == jenis && A[i].judul == judul {
			return i
		}
	}
	return idx
}

func hapus_data() {

}

func tampilkan_data() {

}

func urutkan_data() {

}
