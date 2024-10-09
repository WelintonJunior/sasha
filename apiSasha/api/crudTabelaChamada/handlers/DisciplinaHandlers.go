package handlers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type DisciplinaHandler struct {
	commons.GenericHandler[entity.Disciplina]
}

func NewDisciplinaHandler(genericHandler commons.GenericHandler[entity.Disciplina]) *DisciplinaHandler {
	return &DisciplinaHandler{
		genericHandler,
	}
}
