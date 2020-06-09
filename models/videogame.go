package models

import (
	"time"
)

const(
	CatAventure = iota
	CatFPS
	CatStrategy
	CatChess
)

type VideoGame struct {
	ID          string   `json:"uuid"`
	IDUser     string   `json:"user_name"`
	Title       string   `json:"email"`
	Description string   `json:"pass"`
	Images      []string `json:"fn"`
	Status      int      `json:"ln"`
	Category    int	`json:"category"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
	DeleteDate  time.Time `json:"delete_date"`
}


