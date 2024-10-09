package services

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type DetChamadaService struct {
	commons.GenericService[entity.DetalheDaChamada]
}

func NewDetChamadaService(service commons.GenericService[entity.DetalheDaChamada]) *DetChamadaService {
	return &DetChamadaService{service}
}
