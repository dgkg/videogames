package sqlite

import (
	"time"

	"github.com/dgkg/videogames/models"
	"github.com/google/uuid"
)

func (service *Service) AddVideoGame(v *models.VideoGame) error {
	v.ID = uuid.New().String()
	v.CreateDate = time.Now()
	return service.db.Create(&v).Error
}

func (service *Service) GetVideoGames() (map[string]*models.VideoGame, error) {
	var vgs []models.VideoGame
	service.db.Find(&vgs)
	res := make(map[string]*models.VideoGame)
	for i := 0; i < len(vgs); i++ {
		res[vgs[i].ID] = &vgs[i]
	}
	return res, nil
}

func (service *Service) GetVideoGame(uuid string) (*models.VideoGame, error) {
	var vg models.VideoGame
	return &vg, service.db.Where("id = ?", uuid).First(&vg).Error
}

func (service *Service) UpdateVideoGame(uuid string, update map[string]interface{}) (*models.VideoGame, error) {
	service.db.Model(&models.User{}).Where("id = ?", uuid).Updates(update)
	return service.GetVideoGame(uuid)
}

func (service *Service) DeleteVideoGame(uuid string) error {
	vg := models.VideoGame{
		ID: uuid,
	}
	return service.db.Delete(&vg).Error
}
