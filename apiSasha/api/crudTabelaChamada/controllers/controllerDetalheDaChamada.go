package controllers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/handlers"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

type DetChamadaController struct {
	commons.GenericController[entity.DetalheDaChamada]
}

func NewDetChamadaController(genericController commons.GenericController[entity.DetalheDaChamada]) *DetChamadaController {
	return &DetChamadaController{
		genericController,
	}
}

func StartCrudDetalheDaChamada(server *fiber.App) {
	var detChamada entity.DetalheDaChamada
	baseHandler := commons.NewGenericHandler[entity.DetalheDaChamada]()
	detChamadaHandler := handlers.NewDetChamadaHandler(*baseHandler)

	baseService := commons.NewGenericService[entity.DetalheDaChamada]()
	detChamadaService := services.NewDetChamadaService(*baseService)

	baseController := commons.NewGenericController[entity.DetalheDaChamada]()
	detChamadaController := NewDetChamadaController(*baseController)

	detChamadaController.BuildAllRoutes(server, "detChamada", detChamadaHandler, detChamadaService, detChamada.TableName(), detChamada.GetPrimaryKey())
}
