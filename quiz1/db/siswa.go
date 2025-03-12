package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)


type Siswa struct {
	ID     int    `json:"ID"`
	Nama   string `json:"Nama"`
	Email  string `json:"Email"`
	Alamat string `json:"Alamat"`
}


func CreateSiswa(database *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method tidak diperbolehkan", http.StatusMethodNotAllowed)
		return
	}

	// Ambil data dari form (x-www-form-urlencoded)
	nama := r.FormValue("nama")
	email := r.FormValue("email")
	alamat := r.FormValue("alamat")

	// Validasi data
	if nama == "" || email == "" || alamat == "" {
		http.Error(w, "Semua data harus diisi", http.StatusBadRequest)
		return
	}

	// Menyusun query untuk memasukkan data ke database
	query := "INSERT INTO siswa (Nama, Email, Alamat) VALUES (?, ?, ?)"
	_, err := database.Exec(query, nama, email, alamat)
	if err != nil {
		http.Error(w, "Gagal menambahkan siswa", http.StatusInternalServerError)
		return
	}

	// Mengirimkan response jika data berhasil ditambahkan
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data siswa berhasil ditambahkan!"})
}


func GetAllSiswa(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM siswa")
	if err != nil {
		log.Print("Error Query: ", err) 
		http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	var siswaList []Siswa
	for rows.Next() {
		var s Siswa
		if err := rows.Scan(&s.ID, &s.Nama, &s.Email, &s.Alamat); err != nil {
			log.Print("Error Scan: ", err) 
			http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
			return
		}
		siswaList = append(siswaList, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(siswaList)
}


//----
// 
func UpdateSiswa(database *sql.DB, w http.ResponseWriter, r *http.Request) {
    // Ambil ID dari URL
    id := r.URL.Path[len("/siswa/update/"):]

    // Ambil data dari form (x-www-form-urlencoded)
    nama := r.FormValue("nama")
    email := r.FormValue("email")
    alamat := r.FormValue("alamat")

    // Update hanya kolom yang diisi
    query := "UPDATE siswa SET Nama = COALESCE(NULLIF(?, ''), Nama), Email = COALESCE(NULLIF(?, ''), Email), Alamat = COALESCE(NULLIF(?, ''), Alamat) WHERE id = ?"
    _, err := database.Exec(query, nama, email, alamat, id)
    if err != nil {
        http.Error(w, "Gagal memperbarui siswa", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Data siswa berhasil diperbarui!"})
}


func DeleteSiswa(database *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/siswa/delete/"):] // Ambil ID dari URL
	query := "DELETE FROM siswa WHERE id = ?"
	_, err := database.Exec(query, id)
	if err != nil {
		http.Error(w, "Gagal menghapus siswa", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Data siswa berhasil dihapus!"})
}
