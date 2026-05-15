package Api

import (
	"encoding/json"
	"forum-backend/Server/Database"
	"forum-backend/Server/Tools"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var dataFromJS Database.UserDom

	err := json.NewDecoder(r.Body).Decode(&dataFromJS)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	var messageError string

	// Tester le nom d'Utilisateur
	if dataFromJS.Username == "" {
		messageError = "Nom d'utilisateur vide"
	}

	for _, char := range dataFromJS.Username {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z' || char < '0' || char > '9') {
			messageError = "Nom d'utilisateur invalide"
			break
		}
	}

	if len(dataFromJS.Username) > 20 {
		messageError = "Nom d'utilisateur trop long"
	}

	// Tester le mot de passe

	if dataFromJS.Password == "" {
		messageError = "Mot de passe vide"
	}

	if len(dataFromJS.Password) < 8 {
		messageError = "Mot de passe trop court"
	}

	// Retour si il y a un message d'erreur

	if messageError != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": messageError})
		return
	}

	// Hasher mdp
	hashedPassword, err := Tools.HashPassword(dataFromJS.Password)
	if err != nil {
		http.Error(w, "Erreur (HashPassword) : ", http.StatusInternalServerError)
		return
	}

	newUser := dataFromJS.ToUser()
	newUser.Password = hashedPassword

	// Mettre l'heure js
	newUser.LastConnexion = dataFromJS.LastConnexion

	result := Database.DB.Create(&newUser)
	if result.Error != nil {
		http.Error(w, "Erreur lors de la création", http.StatusInternalServerError)
		return
	}

	messageError = "Utilisateur créé avec succès"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": messageError})
}
