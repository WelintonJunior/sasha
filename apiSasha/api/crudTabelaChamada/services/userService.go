package services

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
	"example.com/apiSasha/database"
)

type UserService struct {
	commons.GenericService[entity.User]
}

func NewUserService(service commons.GenericService[entity.User]) *UserService {
	return &UserService{service}
}

func (s *UserService) GetUserById(userID int) (entity.User, error) {
	var resources entity.User

	err := database.DB.Raw(`
		select * from usu_users where usu_ra = ?;
	`, userID).Scan(&resources).Error

	if err != nil {
		return entity.User{}, err
	}

	return resources, nil
}
