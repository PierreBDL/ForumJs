package Api

import (
	"encoding/json"
	"forum-backend/Server/Database"
	"forum-backend/Server/Tools"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var dataFromJS Database.UserDom

	err := json.NewDecoder(r.Body).Decode(&dataFromJS)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	var messageError string

	// Récup du nom d'utilisateur et du mdp
	var user Database.User
	result := Database.DB.Where("username = ?", dataFromJS.Username).First(&user)

	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Identifiants ou mot de passe incorrect"})
		return
	}

	// Comparer hash et mdp dom
	if !Tools.CheckPasswordHash(dataFromJS.Password, user.Password) {
		messageError = "Nom d'utilisateur ou mot de passe incorrect"
	}

	// Retour si il y a un message d'erreur

	if messageError != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": messageError})
		return
	}

	messageError = "Connexion réussie"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": messageError})
}
