package mock

import (
	"errors"
	"time"

	"github.com/dgkg/videogames/models"
	"github.com/google/uuid"
)

func (m DB) AddUser(u *models.User) error {
	u.ID = uuid.New().String()
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

	u2.FirstName = update["FirstName"].(string)
	u2.LastName = update["LastName
	u2.Email = u.Email
	u2.UpdateDate = time.Now()
	m.Users[uuid] = &u2

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
