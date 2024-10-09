package services

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type AlunoService struct {
	commons.GenericService[entity.Aluno]
}

func NewAlunoService(service commons.GenericService[entity.Aluno]) *AlunoService {
	return &AlunoService{service}
}
