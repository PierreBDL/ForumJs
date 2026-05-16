package Api

import (
	"encoding/json"
	"forum-backend/Server/Database"
	"net/http"
)

type PostResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

type ProfilResponse struct {
	Username      string                  `json:"username"`
	AvatarLink    string                  `json:"avatarLink"`
	Role          string                  `json:"role"`
	CreateAt      string                  `json:"createAt"`
	LastConnexion string                  `json:"lastConnexion"`
	IsOnline      bool                    `json:"isOnline"`
	PostCount     int                     `json:"post_count"`
	CommentCount  int                     `json:"comment_count"`
	LikeCount     int                     `json:"like_count"`
	DislikeCount  int                     `json:"dislike_count"`
	Email         string                  `json:"email"`
	LastPosts     map[string]PostResponse `json:"lastPosts"`
}

func ProfilHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username manquant", http.StatusBadRequest)
		return
	}

	var user Database.User
	err := Database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Utilisateur introuvable"})
		return
	}

	var userPosts []Database.Post
	Database.DB.Where("user_id = ?", user.ID).Order("id DESC").Limit(4).Find(&userPosts)

	lastPostsMap := make(map[string]PostResponse)
	for i, p := range userPosts {
		key := "post" + string(rune(49+i))
		lastPostsMap[key] = PostResponse{
			Title:   p.Title,
			Content: p.Content,
			Date:    p.CreatedAt.Format("2006-01-02"),
		}
	}

	response := ProfilResponse{
		Username:      user.Username,
		AvatarLink:    user.AvatarLink,
		Role:          user.Role,
		CreateAt:      user.CreatedAt.Format("2006-01-02"),
		LastConnexion: user.LastConnexion.Format("2006-01-02"),
		IsOnline:      user.Is_online,
		PostCount:     user.PostsCreated,
		CommentCount:  user.PostsCommented,
		LikeCount:     user.PostsLiked,
		DislikeCount:  user.PostsDisliked,
		Email:         user.Email,
		LastPosts:     lastPostsMap,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
