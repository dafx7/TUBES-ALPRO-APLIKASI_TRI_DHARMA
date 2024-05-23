package main

import "fmt"

const NMAX int = 10

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
	asciiArt := `
 __  __ _____ _   _ _   _ 
|  \/  | ____| \ | | | | |
| |\/| |  _| |  \| | | | |
| |  | | |___| |\  | |_| |
|_|  |_|_____|_| \_|\___/ 
`
	fmt.Println(asciiArt)
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
		fmt.Print("\033[H\033[2J")
		if pilihan == 1 {
			tambah_data(&ArrayPPM, &nPPM)
		} else if pilihan == 2 {
			edit_data(&ArrayPPM, &nPPM)
		} else if pilihan == 3 {
			hapus_data(&ArrayPPM, &nPPM)
		} else if pilihan == 4 {
			tampilkan_data(ArrayPPM, nPPM)
		} else if pilihan == 5 {
			urutkan_data()
		}
		tulisan_menu()
		fmt.Scan(&pilihan)
		fmt.Println()
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
	fmt.Print("\033[H\033[2J")
	fmt.Println("DATA BERHASIL DITAMBAHKAN.")
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
		fmt.Println("Data yang ingin diedit tidak di temukan.")
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

func hapus_data(A *arrPPM, n *int) {
	var pilihan int
	fmt.Print("Input data ke berapa yang ingin di hapus: ")
	fmt.Scan(&pilihan)

	if pilihan > *n {
		fmt.Print("Data yang ingin di hapus tidak ada.")
		return
	}
	fmt.Println(*n)

	// Jika data yang di hapus bukan data yang terakhir atau NMAX maka penghapusan data dilakukan dengan metode menimpa.
	if pilihan != NMAX {
		for i := pilihan - 1; i < *n; i++ {
			A[i] = A[i+1]
		}
	} else {
		// Tetapi jika data yang dihapus adalah data ke NMAX maka data index terakhir harus dijadikan himpunan kosong.
		A[pilihan-1] = PPM{}
	}

	*n--
}

func tampilkan_data(A arrPPM, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("----Data ke-%d----\n", i+1)
		fmt.Printf("Jenis: %s\n", A[i].jenis)
		fmt.Printf("Judul: %s\n", A[i].judul)
		fmt.Printf("Ketua: %s\n", A[i].ketua)
		for j := 0; j < A[i].jumAnggota; j++ {
			fmt.Printf("Anggota ke-%d: %s\n", j+1, A[i].anggota[j])
		}
		fmt.Printf("Prodi/Fakultas: %s\n", A[i].prodi)
		fmt.Printf("Sumber Dana: %s\n", A[i].sumber_dana)
		fmt.Printf("Luaran: %s\n", A[i].luaran)
		fmt.Printf("Tahun kegiatan: %d\n", A[i].tahun_kegiatan)
	}
}

func urutkan_data() {

}
