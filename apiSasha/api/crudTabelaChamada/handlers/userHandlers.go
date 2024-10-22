package handlers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type UserHandler struct {
	commons.GenericHandler[entity.User]
}

func NewUserHandler(genericHandler commons.GenericHandler[entity.User]) *UserHandler {
	return &UserHandler{
		genericHandler,
	}
}
