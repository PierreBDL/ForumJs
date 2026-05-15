package Api

import (
	"encoding/json"
	"forum-backend/Server/Database"
	"net/http"
)

func EditProfilHandler(w http.ResponseWriter, r *http.Request) {
	// Autoriser methode Post
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Struc données Dom
	var input struct {
		LastUsername string `json:"lastUserNamer"`
		Username     string `json:"username"`
		Email        string `json:"email"`
		Avatar       string `json:"avatar"`
	}

	// Récup données
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Mise à jour profil
	var user Database.User
	result := Database.DB.Where("username = ?", input.LastUsername).First(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur introuvable"})
		return
	}

	// Chercher si le pseudo est déjà pris
	if input.Username != input.LastUsername {
		var check Database.User
		err := Database.DB.Where("username = ?", input.Username).First(&check).Error
		if err == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "Ce nom d'utilisateur est déjà pris !"})
			return
		}
	}

	// Vérif si l'adresse courriel est dans un format valide
	isContainArobase := false
	isContainPoint := false
	for _, char := range input.Email {
		if char == '@' {
			isContainArobase = true
		}
		if char == '.' {
			isContainPoint = true
		}
	}

	if !isContainArobase || !isContainPoint {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Adresse courriel invalide"})
		return
	}

	user.Username = input.Username
	user.Email = input.Email
	user.AvatarLink = input.Avatar

	result = Database.DB.Save(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Erreur lors de la mise à jour"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Profil mis à jour avec succès"})
}
