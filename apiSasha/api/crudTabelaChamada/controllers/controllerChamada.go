package controllers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/handlers"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

func StartCrudChamada(server *fiber.App) {
	var chamada entity.Chamada
	baseHandler := commons.NewGenericHandler[entity.Chamada]()
	chamadaHandler := handlers.NewChamadaHandler(*baseHandler)

	baseService := commons.NewGenericService[entity.Chamada]()
	chamadaService := services.NewChamadaService(*baseService)

	baseController := commons.NewGenericController[entity.Chamada]()

	baseController.BuildAllRoutes(server, "chamada", chamadaHandler, chamadaService, chamada.TableName(), chamada.GetPrimaryKey())
}
