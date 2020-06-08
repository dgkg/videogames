package main

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

var Users map[string]User

func init() {
	Users = make(map[string]User)

	var u User
	u.UserName = "Philipe"
	u.Password = "boujour123"

	Users["aoidifhlkajdf"] = u
}

func main() {
	r := gin.Default()
	r.GET("/users/:uuid", func(ctx *gin.Context) {
		ctx.JSON(200, Users[ctx.Param("uuid")])
	})
	r.Run(":8080")
}

// User defin a single User.
type User struct {
	ID          string    `json:"uuid"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Password    string    `json:"pass"`
	FirstName   string    `json:"fn"`
	LastName    string    `json:"ln"`
	DateOfBirth time.Time `json:"date_of_birth"`
	CreateDate  time.Time `json:"create_date"`
	DeleteDate  time.Time `json:"delete_date"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type aux struct {
		ID          string    `json:"uuid"`
		UserName    string    `json:"user_name"`
		Email       string    `json:"email"`
		FirstName   string    `json:"fn"`
		LastName    string    `json:"ln"`
		DateOfBirth time.Time `json:"date_of_birth"`
		CreateDate  time.Time `json:"create_date"`
		DeleteDate  time.Time `json:"delete_date"`
	}
	var a aux
	a.ID = u.ID
	a.UserName = u.UserName
	a.Email = u.Email
	a.FirstName = u.FirstName
	a.LastName = u.LastName
	a.DateOfBirth = u.DateOfBirth
	a.CreateDate = u.CreateDate
	a.DeleteDate = u.DeleteDate

	return json.Marshal(a)
}
