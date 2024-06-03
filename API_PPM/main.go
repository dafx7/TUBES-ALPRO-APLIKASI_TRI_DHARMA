package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Define the PPM struct
type PPM struct {
	ID            int       `json:"id"`
	Jenis         string    `json:"jenis"`
	Ketua         string    `json:"ketua"`
	Prodi         string    `json:"prodi"`
	Judul         string    `json:"judul"`
	SumberDana    string    `json:"sumber_dana"`
	Luaran        string    `json:"luaran"`
	TahunKegiatan int       `json:"tahun_kegiatan"`
	JumAnggota    int       `json:"jumAnggota"`
	Anggota       [4]string `json:"anggota"`
}

var db *sql.DB
var err error

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/tri_dharma"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database successfully.")
}

func getPPMs(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, jenis, ketua, prodi, judul, sumber_dana, luaran, tahun_kegiatan, jumAnggota, anggota1, anggota2, anggota3, anggota4 FROM ppm")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ppms []PPM
	for rows.Next() {
		var ppm PPM
		if err := rows.Scan(&ppm.ID, &ppm.Jenis, &ppm.Ketua, &ppm.Prodi, &ppm.Judul, &ppm.SumberDana, &ppm.Luaran, &ppm.TahunKegiatan, &ppm.JumAnggota, &ppm.Anggota[0], &ppm.Anggota[1], &ppm.Anggota[2], &ppm.Anggota[3]); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ppms = append(ppms, ppm)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ppms)
}

func createPPM(w http.ResponseWriter, r *http.Request) {
	var ppm PPM
	if err := json.NewDecoder(r.Body).Decode(&ppm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO ppm (jenis, ketua, prodi, judul, sumber_dana, luaran, tahun_kegiatan, jumAnggota, anggota1, anggota2, anggota3, anggota4) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, ppm.Jenis, ppm.Ketua, ppm.Prodi, ppm.Judul, ppm.SumberDana, ppm.Luaran, ppm.TahunKegiatan, ppm.JumAnggota, ppm.Anggota[0], ppm.Anggota[1], ppm.Anggota[2], ppm.Anggota[3])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ppm.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ppm)
}

func updatePPM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var ppm PPM
	if err := json.NewDecoder(r.Body).Decode(&ppm); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE ppm SET jenis=?, ketua=?, prodi=?, judul=?, sumber_dana=?, luaran=?, tahun_kegiatan=?, jumAnggota=?, anggota1=?, anggota2=?, anggota3=?, anggota4=? WHERE id=?`
	_, err := db.Exec(query, ppm.Jenis, ppm.Ketua, ppm.Prodi, ppm.Judul, ppm.SumberDana, ppm.Luaran, ppm.TahunKegiatan, ppm.JumAnggota, ppm.Anggota[0], ppm.Anggota[1], ppm.Anggota[2], ppm.Anggota[3], id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ppm)
}

func deletePPM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	query := `DELETE FROM ppm WHERE id=?`
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ppms", getPPMs).Methods("GET")
	r.HandleFunc("/ppms", createPPM).Methods("POST")
	r.HandleFunc("/ppms/{id}", updatePPM).Methods("PUT")
	r.HandleFunc("/ppms/{id}", deletePPM).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
