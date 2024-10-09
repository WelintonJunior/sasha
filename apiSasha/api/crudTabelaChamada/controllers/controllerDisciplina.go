package controllers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/handlers"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

type DisciplinaController struct {
	commons.GenericController[entity.Disciplina]
}

func NewDisciplinaController(genericController commons.GenericController[entity.Disciplina]) *DisciplinaController {
	return &DisciplinaController{
		genericController,
	}
}

func StartCrudDisciplina(server *fiber.App) {
	var disciplina entity.Disciplina
	baseHandler := commons.NewGenericHandler[entity.Disciplina]()
	disciplinaHandler := handlers.NewDisciplinaHandler(*baseHandler)

	baseService := commons.NewGenericService[entity.Disciplina]()
	disciplinaService := services.NewDisciplinaService(*baseService)

	baseController := commons.NewGenericController[entity.Disciplina]()
	disciplinaController := NewDisciplinaController(*baseController)

	disciplinaController.BuildAllRoutes(server, "disciplina", disciplinaHandler, disciplinaService, disciplina.TableName(), disciplina.GetPrimaryKey())
}
