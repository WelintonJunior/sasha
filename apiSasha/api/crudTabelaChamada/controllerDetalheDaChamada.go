package crudTabelaChamada

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

func StartCrudDetalheDaChamada(server *fiber.App) {
	var detChamada entity.DetalheDaChamada
	baseHandler := commons.NewGenericHandler[entity.DetalheDaChamada]()
	baseService := commons.NewGenericService[entity.DetalheDaChamada]()
	baseController := commons.NewGenericController[entity.DetalheDaChamada]()

	baseController.BuildAllRoutes(server, "detChamada", *baseHandler, *baseService, detChamada.TableName())
}
