package main

import (
	"html/template"
	"net/http"
        "os" 
	"fmt" 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/1amkaizen/telegoGPT/models"
)

func displayData(w http.ResponseWriter, r *http.Request) {
	// Menghubungkan ke database
// Konfigurasi koneksi ke database
	// Baca konfigurasi dari environment variables
    // Konfigurasi koneksi ke database
	dsn := "root:MR8MPoeiVJdHcaDrsVjF@tcp(containers-us-west-150.railway.app:5616)/railway"

  
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengambil data dari database
	var messages []models.Messages
	db.Find(&messages)

	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Data from Database</title>
	</head>
	<body>
		<h1>Data from Database</h1>
		<ul>
			{{range .}}
			<li>
				ID: {{.Id}}, Message ID: {{.MessageID}}, User ID: {{.UserID}}, Message: {{.Message}}, Reply: {{.Reply}}, Created At: {{.CreatedAt}}
			</li>
			{{end}}
		</ul>
	</body>
	</html>`

	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", displayData)
	http.ListenAndServe(":8080", nil)
}
