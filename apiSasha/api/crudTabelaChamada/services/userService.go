package services

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type UserService struct {
	commons.GenericService[entity.User]
}

func NewUserService(service commons.GenericService[entity.User]) *UserService {
	return &UserService{service}
}
