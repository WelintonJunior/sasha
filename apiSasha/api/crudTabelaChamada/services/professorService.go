package services

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type ProfessorService struct {
	commons.GenericService[entity.Professor]
}

func NewProfessorService(service commons.GenericService[entity.Professor]) *ProfessorService {
	return &ProfessorService{service}
}
