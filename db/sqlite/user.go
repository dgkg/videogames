package sqlite

func (service *Service) AddUser(u *models.User) error {
	u.ID = uuid.New().String()
	u.CreateDate=  time.Now()
	return service.db.Create(&u).Error
}

func (service *Service) GetUsers() (map[string]*models.User, error) {
	us := make([]models.User)
	service.db.Find(&us)
	res := make(map[string]*models.User)
	for i := 0; i < len(vgs); i++ {
		res[us[i].ID] = &us[i]
	}
	return res, nil
}

func (service *Service) GetUser(uuid string) (*models.User, error) {
	var u models.User
	return vg, service.db.Where("id = ?", uuid).First(&u).Error
}

func (service *Service) UpdateUser(uuid string, update map[string]interface{}) (*models.User, error) {
	service.db.Model(&models.User).Where("id = ?", uuid).Updates(update)
	return nil, nil
}

func (service *Service) DeleteUser(uuid string) error {
	u := models.User{
		ID:uuid,
	}
	return service.db.Delete(&u).Error
}
