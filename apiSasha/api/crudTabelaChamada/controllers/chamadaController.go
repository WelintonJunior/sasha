package controllers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/handlers"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

type ChamadaController struct {
	commons.GenericController[entity.Chamada]
}

func NewChamadaController(genericController commons.GenericController[entity.Chamada]) *ChamadaController {
	return &ChamadaController{
		genericController,
	}
}

func StartCrudChamada(server *fiber.App) {
	var chamada entity.Chamada
	baseHandler := commons.NewGenericHandler[entity.Chamada]()
	chamadaHandler := handlers.NewChamadaHandler(*baseHandler)

	baseService := commons.NewGenericService[entity.Chamada]()
	chamadaService := services.NewChamadaService(*baseService)

	baseController := commons.NewGenericController[entity.Chamada]()
	chamadaController := NewChamadaController(*baseController)

	server.Get("getAllDetChamadaFromChamadaId/:id", func(ctx *fiber.Ctx) error {
		return chamadaHandler.GetAllDetChamadaFromChamadaId(ctx, *chamadaService)
	})

	server.Get("getTokenFromChamadaId/:id", func(ctx *fiber.Ctx) error {
		return chamadaHandler.GetTokenFromChamadaId(ctx, *chamadaService)
	})

	server.Put("updatePresencaAluno/:id", func(ctx *fiber.Ctx) error {
		return chamadaHandler.UpdatePresencaAluno(ctx, *chamadaService)
	})

	server.Put("updateTokenInChamada", func(ctx *fiber.Ctx) error {
		return chamadaHandler.UpdateTokenInChamada(ctx, *chamadaService)
	})

	chamadaController.BuildAllRoutes(server, "chamada", chamadaHandler, chamadaService, chamada.TableName(), chamada.GetPrimaryKey())
}
