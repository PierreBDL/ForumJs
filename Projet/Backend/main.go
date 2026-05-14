package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../Frontend/Templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// Static files
	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("../Frontend/Assets"))))

	http.ListenAndServe(":8080", nil)
}
