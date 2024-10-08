package initApi

import (
	"example.com/apiSasha/api/crudTabelaChamada"
	"github.com/gofiber/fiber/v2"
)

func InitApi(server *fiber.App) {
	crudTabelaChamada.StartCrudChamada(server)
	crudTabelaChamada.StartCrudDetalheDaChamada(server)
	crudTabelaChamada.StartCrudDisciplina(server)
	crudTabelaChamada.StartCrudProfessor(server)
	crudTabelaChamada.StartCrudAluno(server)
}
