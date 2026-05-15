package Database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&User{}, &Post{})

	fillDB()
}

func GetDB() *gorm.DB {
	return DB
}

// Données IA pour tester

func fillDB() {
	var count int64
	DB.Model(&Post{}).Count(&count)

	if count == 0 {
		testUser := User{
			Username:   "Lucas_Dev",
			Email:      "lucas@test.com",
			Password:   "testtest",
			AvatarLink: "/Assets/Images/profil.jpg",
			Role:       "Admin",
		}
		DB.Create(&testUser)

		dummyPosts := []Post{
			{
				Title:   "Développer un routeur en Go sans framework",
				Content: "C'est fou à quel point le package net/http de Go est puissant. Pas besoin de Gin ou de Fiber pour faire une API REST propre, performante et facile à maintenir !",
				UserID:  testUser.ID,
			},
			{
				Title:   "Pourquoi SQLite est incroyable pour les petits projets",
				Content: "Pas besoin de configurer un serveur Docker ou Postgres complet. Un simple fichier local .db et GORM s'occupe de tout. C'est parfait pour prototyper rapidement.",
				UserID:  testUser.ID,
			},
			{
				Title:   "Besoin d'aide : Problème d'asynchronisme en JS",
				Content: "J'essaie de faire un fetch sur mon API Go mais ma fonction d'affichage s'exécute avant que le serveur n'ait répondu... Quelqu'un a une astuce avec async/await ?",
				UserID:  testUser.ID,
			},
		}

		DB.Create(&dummyPosts)
	}
}
