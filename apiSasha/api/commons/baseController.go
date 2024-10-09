package commons

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type BaseController[T any] interface {
	BuildCreateRoute(server *fiber.App, path string, handler BaseHandler[T], service BaseService[T], tableName string) error
	BuildGetRoute(server *fiber.App, path string, handler BaseHandler[T], service BaseService[T], tableName string) error
	BuildDeleteRoute(server *fiber.App, path string, handler BaseHandler[T], service BaseService[T], tableName, primaryKey string) error
	BuildUpdateRoute(server *fiber.App, path string, handler BaseHandler[T], service BaseService[T], tableName, primaryKey string) error
}

type GenericController[T any] struct {
}

func NewGenericController[T any]() *GenericController[T] {
	return &GenericController[T]{}
}

func (c *GenericController[T]) BuildAllRoutes(server *fiber.App, path string, handler BaseHandler[T], service BaseService[T], tableName string, primaryKey string) {
	err := c.BuildCreateRoute(server, fmt.Sprintf("/%s/create", path), handler, service, tableName, primaryKey)
	if err != nil {
		panic(err)
	}

	err = c.BuildGetRoute(server, fmt.Sprintf("/%s", path), handler, service, tableName, primaryKey)
	if err != nil {
		panic(err)
	}

	err = c.BuildDeleteRoute(server, fmt.Sprintf("/%s/:id", path), handler, service, tableName, primaryKey)
	if err != nil {
		panic(err)
	}

	err = c.BuildUpdateRoute(server, fmt.Sprintf("/%s/:id", path), handler, service, tableName, primaryKey)
	if err != nil {
		panic(err)
	}
}

func (c *GenericController[T]) BuildCreateRoute(server *fiber.App, path string, handler BaseHandler[T], service BaseService[T], tableName, primaryKey string) error {
	server.Post(path, func(ctx *fiber.Ctx) error {

		return handler.BuildCreateHandler(ctx, service, tableName)
	})
	return nil
}

func (c *GenericController[T]) BuildGetRoute(server *fiber.App, path string, handler BaseHandler[T], service BaseService[T], tableName, primaryKey string) error {
	server.Get(path, func(ctx *fiber.Ctx) error {
		return handler.BuildGetHandler(ctx, service, tableName)
	})
	return nil
}

func (c *GenericController[T]) BuildDeleteRoute(server *fiber.App, path string, handler BaseHandler[T], service BaseService[T], tableName, primaryKey string) error {
	server.Delete(path, func(ctx *fiber.Ctx) error {
		return handler.BuildDeleteHandler(ctx, service, tableName, primaryKey)
	})

	return nil
}

func (c *GenericController[T]) BuildUpdateRoute(server *fiber.App, path string, handler BaseHandler[T], service BaseService[T], tableName, primaryKey string) error {
	server.Put(path, func(ctx *fiber.Ctx) error {
		return handler.BuildUpdateHandler(ctx, service, tableName, primaryKey)
	})

	return nil
}
