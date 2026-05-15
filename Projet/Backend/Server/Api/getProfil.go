package Api

import (
	"encoding/json"
	"forum-backend/Server/Database"
	"net/http"
)

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	// Autoriser que Get
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Réccup username
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Nom d'utilisateur manquant", http.StatusBadRequest)
		return
	}

	var user Database.User
	result := Database.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur introuvable"})
		return
	}

	profil := user.ToProfil()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profil)
}
