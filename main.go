package main

import (
    "fmt"
    "html/template"
    "net/http"
  "github.com/1amkaizen/telegoGPT/models"
)

func displayData(w http.ResponseWriter, r *http.Request) {
	// Ambil data dari database
	var messages []models.Messages
	models.DB.Find(&messages)

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
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
