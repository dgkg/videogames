package models

import (
	"time"
)

const (
	CatAventure = iota
	CatFPS
	CatStrategy
	CatChess
)

type VideoGame struct {
	ID          string    `json:"uuid"`
	IDUser      string    `json:"user_id"`
	Title       string    `json:"email"`
	Description string    `json:"desc"`
	Images      []string  `json:"imgs"`
	Status      int       `json:"status"`
	Category    int       `json:"category"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_date"`
	DeleteDate  time.Time `json:"delete_date"`
}
