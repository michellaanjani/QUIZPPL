package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"db"
)

func main() {
	// Inisialisasi database
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	defer database.Close()

	r := http.NewServeMux()

	// Menggunakan closure agar handler menerima database instance
	r.HandleFunc("/siswa/create", func(w http.ResponseWriter, r *http.Request) {
		db.CreateSiswa(database, w, r)
	})
	r.HandleFunc("/siswa", func(w http.ResponseWriter, r *http.Request) {
		db.GetAllSiswa(w, r)
	})
	r.HandleFunc("/siswa/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		db.UpdateSiswa(database, w, r)
	})
	
	r.HandleFunc("/siswa/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		db.DeleteSiswa(database, w, r)
	})

	fmt.Println("✅ Server berjalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

