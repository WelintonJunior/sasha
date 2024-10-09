package handlers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type ChamadaHandler struct {
	commons.GenericHandler[entity.Chamada]
}

func NewChamadaHandler(genericHandler commons.GenericHandler[entity.Chamada]) *ChamadaHandler {
	return &ChamadaHandler{
		genericHandler,
	}
}
