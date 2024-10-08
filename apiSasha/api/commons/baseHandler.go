package commons

import (
	"github.com/gofiber/fiber/v2"
)

type BaseHandler[T any] interface {
	BuildCreatehandler(ctx *fiber.Ctx, service GenericService[T]) error
}

type GenericHandler[T any] struct {
}

func NewGenericHandler[T any]() *GenericHandler[T] {
	return &GenericHandler[T]{}
}

func (h *GenericHandler[T]) BuildCreatehandler(ctx *fiber.Ctx, service GenericService[T], tableName string) error {
	var resource T

	if err := ctx.BodyParser(&resource); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Erro ao receber dados"})
		return err
	}

	if err := service.BuildCreateService(resource, tableName); err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": "Sucesso!"})
}
