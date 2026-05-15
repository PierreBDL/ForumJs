package Api

import (
	"encoding/json"
	"forum-backend/Server/Database"
	"net/http"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Autoriser que Get
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récup les posts aléatoirements (10 max)
	var posts []Database.Post

	result := Database.DB.Preload("User").Order("RANDOM()").Limit(10).Find(&posts)
	if result.Error != nil {
		http.Error(w, "Erreur lors de la récupération des posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
