package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	 "time" 
	"strings" 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/1amkaizen/telegoGPT/models"
)


type PageData struct {
	MessageCount int
	Messages     []models.Messages
}


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

// Log informasi permintaan HTTP yang masuk
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
// Log informasi mengenai alamat IP pengirim
	log.Printf("Client IP address: %s", r.RemoteAddr)
// Log jumlah pesan yang ditemukan di database
	log.Printf("Jumlah pesan ditemukan: %d", len(messages))
// Log alamat database yang digunakan
        log.Printf("Database address: %s", dsn)
// Log waktu saat server dimulai
	log.Printf("Server start time: %s", time.Now().Format("2006-01-02 15:04:05"))

	// Di dalam fungsi displayData Anda
err = tmpl.Execute(w, len(messages))
if err != nil {
    http.Error(w, "Template execution error", http.StatusInternalServerError)
    log.Println("Template execution error:", err)
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
	// Log bahwa template telah berhasil dikirim
	log.Println("Template successfully sent")
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
 
