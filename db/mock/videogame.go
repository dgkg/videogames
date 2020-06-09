package mock

import (
	"errors"

	"github.com/google/uuid"

	"github.com/dgkg/videogames/models"
)

// MockVideoGame defines a Mocked videogame.
func MockVideoGame(db *db, v struct {
	title, description string
	status, category   int
	images             []string
}) {

	id := uuid.New().String()

	u := models.VideoGame{
		ID:          id,
		Title:       v.title,
		Description: v.description,
		Images:      v.images,
		Status:      v.status,
		Category:    v.category,
	}

	db.VideoGames[id] = &u
}

func (m DB) AddVideoGame(u *models.VideoGame) error {
	u.ID = uuid.New().String()
	m.VideoGames[u.ID] = u

	return nil
}

func (m DB) GetVideoGames() (map[string]*models.VideoGame, error) {
	return m.VideoGames, nil
}

func (m DB) GetVideoGame(uuid string) (*models.VideoGame, error) {
	u, ok := m.VideoGames[uuid]
	if ok == false {
		return nil, errors.New("db mock: user not found")
	}
	return u, nil
}

func (m DB) UpdateVideoGame(uuid string, mapper map[string]interface{}) (*models.VideoGame, error) {
	_, ok := m.Users[uuid]
	if ok == false {
		return nil, errors.New("db mock: user not found")
	}

	for key, value := range mapper {
		if u2[key] != u[key] {
			u2[key] = u[key]
		}
	}

	m.VideoGames[uuid] = &u2

	return m.VideoGames[uuid], nil
}

func (m DB) DeleteVideoGame(uuid string) error {
	_, ok := m.VideoGames[uuid]
	if ok == false {
		return errors.New("db mock: user not found")
	}

	delete(m.VideoGames, uuid)

	return nil
}
