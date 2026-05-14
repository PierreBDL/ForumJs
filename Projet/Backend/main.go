package main

import (
	"forum-backend/Server/Api"
	"forum-backend/Server/Database"
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
	http.HandleFunc("/register", Frontend.DisplayPages)
	http.HandleFunc("/login", Frontend.DisplayPages)

	// Static files
	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("../Frontend/Assets"))))

	// BDD
	Database.InitDB()

	http.HandleFunc("/api/register", Api.RegisterHandler)
	http.HandleFunc("/api/login", Api.LoginHandler)

	http.ListenAndServe(":8080", nil)

}
