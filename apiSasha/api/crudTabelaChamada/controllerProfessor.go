package crudTabelaChamada

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

func StartCrudProfessor(server *fiber.App) {
	var professor entity.Professor
	baseHandler := commons.NewGenericHandler[entity.Professor]()
	baseService := commons.NewGenericService[entity.Professor]()
	baseController := commons.NewGenericController[entity.Professor]()

	baseController.BuildAllRoutes(server, "professor", *baseHandler, *baseService, professor.TableName())
}
