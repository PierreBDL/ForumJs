package main

import (
	"forum-backend/Server/Api"
	"forum-backend/Server/Database"
	"forum-backend/Server/Frontend"
	"net/http"
)

func main() {

	// API
	http.HandleFunc("/api/register", Api.RegisterHandler)
	http.HandleFunc("/api/login", Api.LoginHandler)
	http.HandleFunc("/api/profil", Api.ProfilHandler)
	http.HandleFunc("/api/editProfile", Api.EditProfilHandler)
	http.HandleFunc("/api/getPosts", Api.GetPostsHandler)

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

	http.ListenAndServe(":8080", nil)

}
