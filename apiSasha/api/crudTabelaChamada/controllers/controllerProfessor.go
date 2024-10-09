package controllers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/handlers"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

type ProfessorController struct {
	commons.GenericController[entity.Professor]
}

func NewProfessorController(genericController commons.GenericController[entity.Professor]) *ProfessorController {
	return &ProfessorController{
		genericController,
	}
}

func StartCrudProfessor(server *fiber.App) {
	var professor entity.Professor
	baseHandler := commons.NewGenericHandler[entity.Professor]()
	professorHandler := handlers.NewProfessorHandler(*baseHandler)

	baseService := commons.NewGenericService[entity.Professor]()
	professorService := services.NewProfessorService(*baseService)

	baseController := commons.NewGenericController[entity.Professor]()
	professorController := NewProfessorController(*baseController)

	professorController.BuildAllRoutes(server, "professor", professorHandler, professorService, professor.TableName(), professor.GetPrimaryKey())
}
