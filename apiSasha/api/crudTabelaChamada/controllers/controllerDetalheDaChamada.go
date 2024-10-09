package controllers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/handlers"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

func StartCrudDetalheDaChamada(server *fiber.App) {
	var detChamada entity.DetalheDaChamada
	baseHandler := commons.NewGenericHandler[entity.DetalheDaChamada]()
	detChamadaHandler := handlers.NewDetChamadaHandler(*baseHandler)

	baseService := commons.NewGenericService[entity.DetalheDaChamada]()
	detChamadaService := services.NewDetChamadaService(*baseService)

	baseController := commons.NewGenericController[entity.DetalheDaChamada]()

	baseController.BuildAllRoutes(server, "detChamada", detChamadaHandler, detChamadaService, detChamada.TableName(), detChamada.GetPrimaryKey())
}
