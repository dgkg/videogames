package mock

import (
	"github.com/dgkg/videogames/db"
	"github.com/dgkg/videogames/models"
)

// DB struct
type DB struct {
	Users      map[string]*models.User
	VideoGames map[string]*models.VideoGame
}

// New init
func New() db.Store {
	return &DB{
		Users:      make(map[string]*models.User),
		VideoGames: make(map[string]*models.VideoGame),
	}
}
