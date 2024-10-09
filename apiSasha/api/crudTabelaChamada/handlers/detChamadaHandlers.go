package handlers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
)

type DetChamadaHandler struct {
	commons.GenericHandler[entity.DetalheDaChamada]
}

func NewDetChamadaHandler(genericHandler commons.GenericHandler[entity.DetalheDaChamada]) *DetChamadaHandler {
	return &DetChamadaHandler{
		genericHandler,
	}
}
