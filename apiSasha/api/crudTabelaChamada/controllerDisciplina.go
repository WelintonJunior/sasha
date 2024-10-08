package crudTabelaChamada

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

func StartCrudDisciplina(server *fiber.App) {
	var disciplina entity.Disciplina
	baseHandler := commons.NewGenericHandler[entity.Disciplina]()
	baseService := commons.NewGenericService[entity.Disciplina]()
	baseController := commons.NewGenericController[entity.Disciplina]()

	baseController.BuildAllRoutes(server, "disciplina", *baseHandler, *baseService, disciplina.TableName())
}
