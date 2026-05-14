package main

import (
	"forum-backend/Server/Frontend"
	"net/http"
)

func main() {
	// Routes
	http.HandleFunc("/", Frontend.DisplayPages)
	http.HandleFunc("/home", Frontend.DisplayPages)
	http.HandleFunc("/discussions", Frontend.DisplayPages)
	http.HandleFunc("/profil", Frontend.DisplayPages)
	http.HandleFunc("/connexion", Frontend.DisplayPages)
	http.HandleFunc("/inscription", Frontend.DisplayPages)

	// Static files
	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("../Frontend/Assets"))))

	http.ListenAndServe(":8080", nil)
}
