package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/dgkg/videogames/db"
	"github.com/dgkg/videogames/models"
)

// Service is the SQLite connextion for the persistency.
type Service struct {
	db *gorm.DB
}

// New is creating a new SQLite connection.
func New(fileName string) db.Store {
	db, err := gorm.Open("sqlite3", fileName)
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.VideoGame{})

	return &Service{
		db: db,
	}
}
