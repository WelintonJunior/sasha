package commons

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type BaseController[T any] interface {
	BuildCreateRoute(server *fiber.App, path string, handler GenericHandler[T], service GenericService[T], tableName string) error
	BuildGetRoute(server *fiber.App, path string, handler GenericHandler[T], service GenericService[T], tableName string) error
}

type GenericController[T any] struct {
}

func NewGenericController[T any]() *GenericController[T] {
	return &GenericController[T]{}
}

func (c *GenericController[T]) BuildAllRoutes(server *fiber.App, path string, handler GenericHandler[T], service GenericService[T], tableName string) {
	err := c.BuildCreateRoute(server, fmt.Sprintf("/%s/create", path), handler, service, tableName)
	if err != nil {
		panic(err)
	}

	err = c.BuildGetRoute(server, fmt.Sprintf("/%s/", path), handler, service, tableName)
	if err != nil {
		panic(err)
	}
}

func (c *GenericController[T]) BuildCreateRoute(server *fiber.App, path string, handler GenericHandler[T], service GenericService[T], tableName string) error {
	server.Post(path, func(ctx *fiber.Ctx) error {
		return handler.BuildCreatehandler(ctx, service, tableName)
	})
	return nil
}

func (c *GenericController[T]) BuildGetRoute(server *fiber.App, path string, handler GenericHandler[T], service GenericService[T], tableName string) error {
	server.Get(path, func(ctx *fiber.Ctx) error {
		return handler.BuildGethandler(ctx, service, tableName)
	})
	return nil
}
