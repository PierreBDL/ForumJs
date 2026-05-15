package Database

// Models qui viennent de js pour les utilisateurs

type UserDom struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Avatar        string `json:"avatar"`
	CreateAt      string `json:"createAt"`
	LastConnexion string `json:"lastConnexion"`
}

// Découpage de la requete json

func (u *UserDom) ToUser() User {
	return User{
		Username:   u.Username,
		Email:      u.Email,
		Password:   u.Password,
		AvatarLink: u.Avatar,
	}

}

// Mettre en json

type Profil struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	AvatarLink     string `json:"avatarLink"`
	Role           string `json:"role"`
	CreateAt       string `json:"createAt"`
	LastConnexion  string `json:"lastConnexion"`
	PostsLiked     int    `json:"postsLiked"`
	PostsDisliked  int    `json:"postsDisliked"`
	PostsCreated   int    `json:"postsCreated"`
	PostsCommented int    `json:"postsCommented"`
	IsOnline       bool   `json:"isOnline"`
}

func (u *User) ToProfil() Profil {
	return Profil{
		Username:       u.Username,
		Email:          u.Email,
		AvatarLink:     u.AvatarLink,
		Role:           u.Role,
		LastConnexion:  u.LastConnexion,
		PostsLiked:     u.PostsLiked,
		PostsDisliked:  u.PostsDisliked,
		PostsCreated:   u.PostsCreated,
		PostsCommented: u.PostsCommented,
		IsOnline:       u.Is_online,
		CreateAt:       u.Created_at,
	}
}

// Models qui viennent de js pour les posts

type PostDom struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p *PostDom) ToPost() Post {
	return Post{
		Title:   p.Title,
		Content: p.Content,
	}

}
