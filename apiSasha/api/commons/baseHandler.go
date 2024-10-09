package commons

import (
	"github.com/gofiber/fiber/v2"
)

type BaseHandler[T any] interface {
	BuildCreateHandler(ctx *fiber.Ctx, service BaseService[T], tableName string) error
	BuildGetHandler(ctx *fiber.Ctx, service BaseService[T], tableName string) error
	BuildDeleteHandler(ctx *fiber.Ctx, service BaseService[T], tableName, primaryKey string) error
	BuildUpdateHandler(ctx *fiber.Ctx, service BaseService[T], tableName, primaryKey string) error
}

type GenericHandler[T any] struct {
}

func NewGenericHandler[T any]() *GenericHandler[T] {
	return &GenericHandler[T]{}
}

func (h *GenericHandler[T]) BuildCreateHandler(ctx *fiber.Ctx, service BaseService[T], tableName string) error {
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

func (h *GenericHandler[T]) BuildGetHandler(ctx *fiber.Ctx, service BaseService[T], tableName string) error {

	result, err := service.BuildGetService(tableName)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)
}

func (h *GenericHandler[T]) BuildDeleteHandler(ctx *fiber.Ctx, service BaseService[T], tableName, primaryKey string) error {

	id := ctx.Params("id")

	err := service.BuildDeleteService(id, tableName, primaryKey)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": "Sucesso!"})
}

func (h *GenericHandler[T]) BuildUpdateHandler(ctx *fiber.Ctx, service BaseService[T], tableName, primaryKey string) error {

	id := ctx.Params("id")
	var resource T

	if err := ctx.BodyParser(&resource); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}

	err := service.BuildUpdateService(resource, id, tableName, primaryKey)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": "Sucesso!"})
}
