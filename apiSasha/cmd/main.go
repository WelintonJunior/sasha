package main

import (
	"example.com/apiSasha/database"
	"example.com/apiSasha/initApi"
	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()
	database.InitDB()
	initApi.InitApi(server)
	err := server.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
