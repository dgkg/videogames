package mock

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/dgkg/videogames/models"
)

// MockVideoGame defines a Mocked videogame.
func MockVideoGame(db *DB, title,
	description string,
	status, category int,
	images []string) {

	id := uuid.New().String()

	u := models.VideoGame{
		ID:          id,
		Title:       title,
		Description: description,
		Images:      images,
		Status:      status,
		Category:    category,
	}

	db.VideoGames[id] = &u
}

func (m DB) AddVideoGame(u *models.VideoGame) error {
	u.ID = uuid.New().String()
	u.CreateDate = time.Now()
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
	u, ok := m.VideoGames[uuid]
	if ok == false {
		return nil, errors.New("db mock: user not found")
	}

	u.IDUser = mapper["IDUser"].(string)
	u.Title = mapper["Title"].(string)
	u.Description = mapper["Description"].(string)
	u.Images = mapper["Images"].([]string)
	u.Status = mapper["Status"].(int)
	u.Category = mapper["Category"].(int)

	m.VideoGames[uuid] = u

	return u, nil
}

func (m DB) DeleteVideoGame(uuid string) error {
	_, ok := m.VideoGames[uuid]
	if ok == false {
		return errors.New("db mock: user not found")
	}

	delete(m.VideoGames, uuid)

	return nil
}
