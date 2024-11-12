package handlers

import (
	"example.com/apiSasha/api/commons"
	"example.com/apiSasha/api/crudTabelaChamada/services"
	"example.com/apiSasha/api/entity"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserHandler struct {
	commons.GenericHandler[entity.User]
}

func NewUserHandler(genericHandler commons.GenericHandler[entity.User]) *UserHandler {
	return &UserHandler{
		genericHandler,
	}
}

func (h *UserHandler) GetUserById(ctx *fiber.Ctx, service services.UserService) error {
	userID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Invalid ID"})
	}

	result, err := service.GetUserById(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
