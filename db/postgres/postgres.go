package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/dgkg/videogames/db"
	"github.com/dgkg/videogames/db/sqlite"
)

type Service = sqlite.Service

type Credential struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func (c *Credential) String() string {
	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v ssh=disabled",
		c.Host, c.Port, c.User, c.DBName, c.Password)
}

func New(c *Credential) db.Store {
	db, err := gorm.Open("postgres", c.String())
	if err != nil {
		panic(err)
	}
	return &Service{
		DB: db,
	}
}
