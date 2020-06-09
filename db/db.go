package db

import (
	"github.com/dgkg/videogames/models"
)

type Store interface {
	StoreUser
	StoreVideoGames
}

type StoreUser interface {
	AddUser(u *models.User) error
	GetUsers() (map[string]*models.User, error)
	GetUser(uuid string) (*models.User, error)
	UpdateUser(uuid string, update map[string]interface{}) (*models.User, error)
	DeleteUser(uuid string) error
}

type StoreVideoGames interface {
	AddVideoGame(u *models.VideoGame) error
	GetVideoGames() (map[string]*models.VideoGame, error)
	GetVideoGame(uuid string) (*models.VideoGame, error)
	UpdateVideoGame(uuid string, update map[string]interface{}) (*models.VideoGame, error)
	DeleteVideoGame(uuid string) error
}
