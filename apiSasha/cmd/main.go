package main

import (
	"example.com/apiSasha/database"
	"example.com/apiSasha/initApi"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	server := fiber.New()

	server.Use(cors.New())

	server.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
	}))

	database.InitDB()
	initApi.InitApi(server)

	err := server.Listen(":3000")
	if err != nil {
		log.Panic(err)
	}
}
