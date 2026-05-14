package Database

// Models qui viennent de js pour les utilisateurs

type UserDom struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
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
