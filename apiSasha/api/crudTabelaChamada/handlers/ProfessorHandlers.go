package handlers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type ProfessorHandler struct {
	commons.GenericHandler[entity.Professor]
}

func NewProfessorHandler(genericHandler commons.GenericHandler[entity.Professor]) *ProfessorHandler {
	return &ProfessorHandler{
		genericHandler,
	}
}
