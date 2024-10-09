package services

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type DisciplinaService struct {
	commons.GenericService[entity.Disciplina]
}

func NewDisciplinaService(service commons.GenericService[entity.Disciplina]) *DisciplinaService {
	return &DisciplinaService{service}
}
