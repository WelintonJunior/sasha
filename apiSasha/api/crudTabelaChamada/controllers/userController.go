package controllers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/handlers"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	commons.GenericController[entity.User]
}

func NewUserController(genericController commons.GenericController[entity.User]) *UserController {
	return &UserController{
		genericController,
	}
}

func StartCrudUser(server *fiber.App) {
	var user entity.User
	baseHandler := commons.NewGenericHandler[entity.User]()
	userHandler := handlers.NewUserHandler(*baseHandler)

	baseService := commons.NewGenericService[entity.User]()
	userService := services.NewUserService(*baseService)

	baseController := commons.NewGenericController[entity.User]()
	userController := NewUserController(*baseController)

	server.Get("user/:id", func(ctx *fiber.Ctx) error {
		return userHandler.GetUserById(ctx, *userService)
	})

	userController.BuildAllRoutes(server, "user", userHandler, userService, user.TableName(), user.GetPrimaryKey())
}
