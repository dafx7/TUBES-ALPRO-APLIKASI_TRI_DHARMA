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
			urutkan_data(&ArrayPPM, nPPM)
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

func hapus_data(A *arrPPM, n *int) {
	var idx int
	var judul string
	fmt.Print("Input judul data yang ingin di hapus: ")
	fmt.Scan(&judul)

	sort_by_judul(&*A, *n)
	idx = binary_search(*A, *n, judul)
	fmt.Println("idx: ", idx)

	if idx != -1 {
		// Jika data yang di hapus bukan data yang terakhir atau NMAX maka penghapusan data dilakukan dengan metode menimpa.
		if idx+1 != NMAX {
			for i := idx; i < *n-1; i++ {
				A[i] = A[i+1]
			}
		} else {
			// Tetapi jika data yang dihapus adalah data ke NMAX maka data index terakhir harus dijadikan himpunan kosong.
			A[idx] = PPM{}
		}

		*n--
	} else {
		fmt.Print("Tidak terdapat data yang memuat judul tersebut.")
	}
}

func cetak_data(A arrPPM, idx int) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("----Data ke-%d----\n", idx+1)
	fmt.Printf("Jenis: %s\n", A[idx].jenis)
	fmt.Printf("Judul: %s\n", A[idx].judul)
	fmt.Printf("Ketua: %s\n", A[idx].ketua)
	for j := 0; j < A[idx].jumAnggota; j++ {
		// Menampilkan anggota ke-j dari data ke-i
		fmt.Printf("Anggota ke-%d: %s\n", j+1, A[idx].anggota[j])
	}
	fmt.Printf("Prodi/Fakultas: %s\n", A[idx].prodi)
	fmt.Printf("Sumber Dana: %s\n", A[idx].sumber_dana)
	fmt.Printf("Luaran: %s\n", A[idx].luaran)
	fmt.Printf("Tahun kegiatan: %d\n", A[idx].tahun_kegiatan)
}

func tampilkan_data(A arrPPM, n int) {
	var jenis, prodi string
	var tahun, pilihan int
	var found bool = false

	// Jika tidak ada di dalam array maka beri tahu pengguna
	if n == 0 {
		fmt.Println("WARNING!")
		fmt.Println("TIDAK ADA DATA YANG DIINPUT.")
	} else {
		// Menampilkan opsi kepada pengguna
		fmt.Println("Opsi Penampilan")
		fmt.Println("1. Tampilkan semua data")
		fmt.Println("2. Tampilkan dengan filter")
		fmt.Print("Input Pilihan: ")
		fmt.Scan(&pilihan) // Membaca pilihan pengguna

		if pilihan == 1 {
			// Jika pilihan 1, tampilkan semua data
			fmt.Println("List Data:")
			for i := 0; i < n; i++ {
				// Menampilkan data ke-i
				cetak_data(A, i)
			}
		} else {
			// Jika pilihan 2, minta input filter dari pengguna
			fmt.Println("Input filter untuk menampilkan data:")
			fmt.Print("Jenis (Penelitian/Abdimas): ")
			fmt.Scan(&jenis) // Membaca filter jenis
			fmt.Print("Tahun: ")
			fmt.Scan(&tahun) // Membaca filter tahun
			fmt.Print("Fakultas/Prodi: ")
			fmt.Scan(&prodi) // Membaca filter prodi

			// Menampilkan data yang sesuai dengan filter
			fmt.Println("List Data:")
			for i := 0; i < n; i++ {
				// Cek apakah data ke-i sesuai dengan filter
				if A[i].jenis == jenis && A[i].tahun_kegiatan == tahun && A[i].prodi == prodi {
					// Satu atau lebih data sesuai filter ditemukan
					found = true
					// Menampilkan data ke-i jika sesuai filter
					cetak_data(A, i)
				}
			}
			// Jika tidak ada yang di temukan sesuai filter maka beri tahu pengguna.
			if !found {
				fmt.Println("Tidak ada data dengan filter yang di input.")
			}
		}
	}
}

func binary_search(A arrPPM, n int, judul string) int {
	var kiri, kanan, mid, found int
	kiri = 0
	kanan = n - 1
	found = -1
	for kiri <= kanan && found == -1 {
		mid = (kiri + kanan) / 2
		if A[mid].judul[0] > judul[0] {
			kanan = mid - 1
		} else if A[mid].judul[0] < judul[0] {
			kiri = mid + 1
		} else {
			return mid
		}
		fmt.Println(A[mid].judul)
	}
	return found
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

func urutkan_data(A *arrPPM, n int) {
	var pilihan int
	fmt.Println("Pilih algoritma yang di pakai untuk sorting:")
	fmt.Println("1. Decending")
	fmt.Println("2. Ascending")

	fmt.Print("Input Pilihan: ")
	fmt.Scan(&pilihan)

	for pilihan != 1 && pilihan != 2 {
		fmt.Println("Error! Coba lagi.")
		fmt.Print("Input Pilihan: ")
		fmt.Scan(&pilihan)
	}

	if pilihan == 1 {
		insertion_sort(&*A, n)
	} else {
		selection_sort(&*A, n)
	}
}

func insertion_sort(A *arrPPM, n int) {
	var tmp PPM
	var j int
	for i := 1; i < n; i++ {
		tmp = A[i]
		j = i - 1
		for j >= 0 && A[j].tahun_kegiatan < tmp.tahun_kegiatan {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = tmp
	}
}

func selection_sort(A *arrPPM, n int) {
	var idx_min int
	var tmp PPM
	for i := 0; i < n; i++ {
		idx_min = i
		// fmt.Println("idxmin: ", idx_min)
		for j := i + 1; j < n; j++ {
			if A[j].tahun_kegiatan < A[idx_min].tahun_kegiatan {
				idx_min = j
			}
		}
		if idx_min != i {
			tmp = A[idx_min]
			A[idx_min] = A[i]
			A[i] = tmp
		}
	}

}

func sort_by_judul(A *arrPPM, n int) {
	var tmp PPM
	var j int
	for i := 1; i < n; i++ {
		tmp = A[i]
		j = i - 1
		for j >= 0 && A[j].judul[0] > tmp.judul[0] {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = tmp
	}
}
