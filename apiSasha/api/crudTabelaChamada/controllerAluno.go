package crudTabelaChamada

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

func StartCrudAluno(server *fiber.App) {
	var aluno entity.Aluno
	baseHandler := commons.NewGenericHandler[entity.Aluno]()
	baseService := commons.NewGenericService[entity.Aluno]()
	baseController := commons.NewGenericController[entity.Aluno]()
	baseController.BuildAllRoutes(server, "aluno", *baseHandler, *baseService, aluno.TableName())
}
