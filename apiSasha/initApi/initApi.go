package initApi

import (
	"example.com/apiSasha/api/crudTabelaChamada/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitApi(server *fiber.App) {
	controllers.StartCrudChamada(server)
	controllers.StartCrudDetalheDaChamada(server)
	controllers.StartCrudDisciplina(server)
	controllers.StartCrudProfessor(server)
	controllers.StartCrudAluno(server)
}
