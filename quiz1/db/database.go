package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)



var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/sekolahdasar"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database: %v", err)
	}

	fmt.Println("✅ Terhubung ke database sekolahdasar")
	return db, nil
}
