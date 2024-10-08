package crudTabelaChamada

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

func StartCrudChamada(server *fiber.App) {
	var chamada entity.Chamada
	baseHandler := commons.NewGenericHandler[entity.Chamada]()
	baseService := commons.NewGenericService[entity.Chamada]()
	baseController := commons.NewGenericController[entity.Chamada]()

	baseController.BuildAllRoutes(server, "chamada", *baseHandler, *baseService, chamada.TableName())
}
