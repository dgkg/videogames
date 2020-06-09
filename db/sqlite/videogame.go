package sqlite

func (serivce *Service) AddVideoGame(v *models.VideoGame) error {
	v.ID = uuid.New().String()
	v.CreateDate=  time.Now()
	return service.db.Create(&v).Error
}

func (serivce *Service) GetVideoGames() (map[string]*models.VideoGame, error) {
	vgs := make([]models.VideoGame)
	service.db.Find(&vgs)
	res := make(map[string]*models.VideoGame)
	for i := 0; i < len(vgs); i++ {
		res[vgs[i].ID] = &vgs[i]
	}
	return res, nil
}

func (serivce *Service) GetVideoGame(uuid string) (*models.VideoGame, error) {
	var vg models.VideoGame
	return vg, service.db.Where("id = ?", uuid).First(&vg).Error
}

func (serivce *Service) UpdateVideoGame(uuid string, update *models.VideoGame) (*models.VideoGame, error) {
	service.db.Model(&models.User).Where("id = ?", uuid).Updates(update)
	return nil, nil
}

func (serivce *Service) DeleteVideoGame(uuid string) error {
	vg := models.VideoGame{
		ID:uuid,
	}
	return service.db.Delete(&vg)
}
