package Database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	Email          string `gorm:"unique;not null"`
	Password       string `gorm:"not null"`
	AvatarLink     string `gorm:"not null"`
	Role           string `gorm:"default:'User'"`
	Created_at     string `gorm:"default:CURRENT_TIMESTAMP"`
	LastConnexion  string `gorm:"default:CURRENT_TIMESTAMP"`
	PostsLiked     int    `gorm:"default:0"`
	PostsDisliked  int    `gorm:"default:0"`
	PostsCreated   int    `gorm:"default:0"`
	PostsCommented int    `gorm:"default:0"`
	Is_online      bool   `gorm:"default:true"`
}

type Post struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Content    string `gorm:"not null"`
	Likes      int    `gorm:"default:0"`
	Dislikes   int    `gorm:"default:0"`
	Comments   int    `gorm:"default:0"`
	Views      int    `gorm:"default:0"`
	Created_at string `gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at string `gorm:"default:CURRENT_TIMESTAMP"`
	UserID     uint   `gorm:"not null"`
	User       User   `gorm:"foreignKey:UserID"`
}
