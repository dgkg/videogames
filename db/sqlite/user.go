package sqlite

import (
	"time"

	"github.com/dgkg/videogames/models"
	"github.com/google/uuid"
)

func (service *Service) AddUser(u *models.User) error {
	u.ID = uuid.New().String()
	u.CreateDate = time.Now()
	return service.db.Create(&u).Error
}

func (service *Service) GetUsers() (map[string]*models.User, error) {
	var us []models.User
	service.db.Find(&us)
	res := make(map[string]*models.User)
	for i := 0; i < len(us); i++ {
		res[us[i].ID] = &us[i]
	}
	return res, nil
}

func (service *Service) GetUser(uuid string) (*models.User, error) {
	var u models.User
	return &u, service.db.Where("id = ?", uuid).First(&u).Error
}

func (service *Service) UpdateUser(uuid string, update map[string]interface{}) (*models.User, error) {
	service.db.Model(&models.User{}).Where("id = ?", uuid).Updates(update)
	return service.GetUser(uuid)
}

func (service *Service) DeleteUser(uuid string) error {
	u := models.User{
		ID: uuid,
	}
	return service.db.Delete(&u).Error
}
