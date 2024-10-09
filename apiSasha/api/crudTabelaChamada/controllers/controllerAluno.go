package controllers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/handlers"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

type AlunoController struct {
	commons.GenericController[entity.Aluno]
}

func NewAlunoController(genericController commons.GenericController[entity.Aluno]) *AlunoController {
	return &AlunoController{
		genericController,
	}
}

func StartCrudAluno(server *fiber.App) {
	var aluno entity.Aluno
	baseHandler := commons.NewGenericHandler[entity.Aluno]()
	alunoHandler := handlers.NewAlunoHandler(*baseHandler)

	baseService := commons.NewGenericService[entity.Aluno]()
	alunoService := services.NewAlunoService(*baseService)

	baseController := commons.NewGenericController[entity.Aluno]()
	alunoController := NewAlunoController(*baseController)

	alunoController.BuildAllRoutes(server, "aluno", alunoHandler, alunoService, aluno.TableName(), aluno.GetPrimaryKey())
}
