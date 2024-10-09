package handlers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type AlunoHandler struct {
	commons.GenericHandler[entity.Aluno]
}

func NewAlunoHandler(genericHandler commons.GenericHandler[entity.Aluno]) *AlunoHandler {
	return &AlunoHandler{
		genericHandler,
	}
}
