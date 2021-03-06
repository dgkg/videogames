package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// User define a single User.
type User struct {
	ID          string    `json:"uuid"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Password    string    `json:"pass"`
	FirstName   string    `json:"fn"`
	LastName    string    `json:"ln"`
	DateOfBirth time.Time `json:"date_of_birth"`
	UpdateDate  time.Time `json:"update_date"`
	CreateDate  time.Time `json:"create_date"`
	DeleteDate  time.Time `json:"delete_date"`
}

func NewUser(name, fn, ln, email, pass string, dateofBirth time.Time) *User {
	return &User{
		UserName:    name,
		FirstName:   fn,
		LastName:    ln,
		Email:       email,
		Password:    pass,
		DateOfBirth: dateofBirth,
		CreateDate:  time.Now(),
	}
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

func (u *User) UnmarshalJSON(data []byte) error {
	type aux struct {
		UserName    string    `json:"user_name"`
		Email       string    `json:"email"`
		FirstName   string    `json:"fn"`
		LastName    string    `json:"ln"`
		Password    string    `json:"pass"`
		DateOfBirth time.Time `json:"date_of_birth"`
		CreateDate  time.Time `json:"create_date"`
		DeleteDate  time.Time `json:"delete_date"`
	}

	var a aux
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	sum := sha256.Sum256([]byte(a.Password))

	u.UserName = a.UserName
	u.Email = a.Email
	u.FirstName = a.FirstName
	u.LastName = a.LastName
	u.Password = fmt.Sprintf("%x", sum)
	u.DateOfBirth = a.DateOfBirth
	u.CreateDate = a.CreateDate
	u.DeleteDate = a.DeleteDate

	return nil
}
