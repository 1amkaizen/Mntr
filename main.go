package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/1amkaizen/telegoGPT/models"
)

func displayData(w http.ResponseWriter, r *http.Request) {
	// Baca konfigurasi dari environment variables
	dsn := "root:5YUTSCHvx0yXwcsJUQFW@tcp(containers-us-west-139.railway.app:6522)/railway" 
	

	// Konfigurasi koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Database connection error:", err)
		return
	}

	// Mengambil data dari database
	var messages []models.Messages
	if err := db.Find(&messages).Error; err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		log.Println("Database query error:", err)
		return
	}

	// Membaca konten dari file index.html
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		log.Println("Template error:", err)
		return
	}

	err = tmpl.Execute(w, messages)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", displayData)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port default jika tidak ada PORT di environment variables
	}
	log.Printf("Server started at :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
 
