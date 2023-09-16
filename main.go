package main

import (
	"html/template"
	"net/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log" 
	"github.com/1amkaizen/telegoGPT/models"
)

func displayData(w http.ResponseWriter, r *http.Request) {
// Baca konfigurasi dari environment variables
// Konfigurasi koneksi ke database
	dsn := "root:5YUTSCHvx0yXwcsJUQFW@tcp(containers-us-west-139.railway.app:6522)/railway" 
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
    log.Fatal("Failed to connect to database:", err)
}
defer db.Close()

log.Println("Connected to database!")



	// Mengambil data dari database
		// Mengambil data dari database
	var messages []models.Messages
	db.Find(&messages)
log.Println("Messages from database:", messages)
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
	http.ListenAndServe(":6522", nil)
	
}
