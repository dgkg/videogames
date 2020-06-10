package mock

import (
	"errors"
	"time"

	"github.com/dgkg/videogames/models"
	"github.com/google/uuid"
)

func MockUser(db *DB, fn, ln, email, pass string) {
	uuid := uuid.New().String()
	db.Users[uuid] = &models.User{
		ID:        uuid,
		FirstName: fn,
		LastName:  ln,
		Email:     email,
		Password:  pass,
	}
}

func (m DB) AddUser(u *models.User) error {
	u.ID = uuid.New().String()
	u.CreateDate = time.Now()
	m.Users[u.ID] = u

	return nil
}

func (m DB) GetUser(uuid string) (*models.User, error) {
	_, ok := m.Users[uuid]
	if ok == false {
		return nil, errors.New("db mock: user not found")
	}

	return m.Users[uuid], nil
}

func (m DB) GetUsers() (map[string]*models.User, error) {
	return m.Users, nil
}

func (m DB) UpdateUser(uuid string, update map[string]interface{}) (*models.User, error) {
	u2, ok := m.Users[uuid]
	if ok == false {
		return nil, errors.New("db mock: user not found")
	}

	if update == nil {
		return nil, errors.New("db mock: the given interface is nil")
	}

	if v, ok := update["fn"]; ok {
		u2.FirstName = v.(string)
	}

	if v, ok := update["ln"]; ok {
		u2.LastName = v.(string)
	}

	if v, ok := update["email"]; ok {
		u2.Email = v.(string)
	}
	u2.UpdateDate = time.Now()
	m.Users[uuid] = u2

	return m.Users[uuid], nil
}

func (m DB) DeleteUser(uuid string) error {
	_, ok := m.Users[uuid]
	if ok == false {
		return errors.New("db mock: user not found")
	}

	delete(m.Users, uuid)

	return nil
}
