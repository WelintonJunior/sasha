package services

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type ChamadaService struct {
	commons.GenericService[entity.Chamada]
}

func NewChamadaService(service commons.GenericService[entity.Chamada]) *ChamadaService {
	return &ChamadaService{service}
}
