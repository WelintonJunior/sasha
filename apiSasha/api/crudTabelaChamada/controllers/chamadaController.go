package controllers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/handlers"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"example.com/apiSasha/api/shared"
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

	server.Get("getChamadaByID/:id", func(ctx *fiber.Ctx) error {
		return chamadaHandler.GetChamadaByID(ctx, *chamadaService)
	})

	server.Get("getAllDetChamadaFromChamadaId/:id", func(ctx *fiber.Ctx) error {
		return chamadaHandler.GetAllDetChamadaFromChamadaId(ctx, *chamadaService)
	})

	server.Get("getTokenFromChamadaId/:id", func(ctx *fiber.Ctx) error {
		return chamadaHandler.GetTokenFromChamadaId(ctx, *chamadaService)
	})

	server.Put("updatePresencaAluno/:id", func(ctx *fiber.Ctx) error {
		err := chamadaHandler.UpdatePresencaAluno(ctx, *chamadaService)
		if err == nil {
			alunoID := ctx.Params("id")
			shared.Broadcast <- "Aluno " + alunoID + " marcou presença"
		}
		return err
	})

	server.Put("updateTokenInChamada", func(ctx *fiber.Ctx) error {
		return chamadaHandler.UpdateTokenInChamada(ctx, *chamadaService)
	})

	server.Put("updateChamadaTrue/:id", func(ctx *fiber.Ctx) error {
		return chamadaHandler.UpdateChamadaTrue(ctx, *chamadaService)
	})

	server.Put("updateChamadaFalse/:id", func(ctx *fiber.Ctx) error {
		return chamadaHandler.UpdateChamadaFalse(ctx, *chamadaService)
	})

	chamadaController.BuildAllRoutes(server, "chamada", chamadaHandler, chamadaService, chamada.TableName(), chamada.GetPrimaryKey())
}
