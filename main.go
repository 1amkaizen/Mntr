package main

import (
	"html/template"
	"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/1amkaizen/telegoGPT/models"
)

func displayData(w http.ResponseWriter, r *http.Request) {
// Baca konfigurasi dari environment variables
// Konfigurasi koneksi ke database
	dsn := "root:MR8MPoeiVJdHcaDrsVjF@tcp(containers-us-west-150.railway.app:5616)/railway"
 
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengambil data dari database
		// Mengambil data dari database
	var messages []models.Messages
	db.Find(&messages)

	// Membaca konten dari file index.html
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}

func main() {
	http.HandleFunc("/", displayData)
	http.ListenAndServe(":8080", nil)
}
