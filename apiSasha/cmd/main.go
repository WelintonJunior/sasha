package main

import (
	"example.com/apiSasha/api/shared"
	"example.com/apiSasha/database"
	"example.com/apiSasha/initApi"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	"log"
	"sync"
)

var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

func main() {
	database.InitDB()
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	initApi.InitApi(app)

	app.Get("/ws", websocket.New(handleConnections))

	go handleMessages()

	log.Fatal(app.Listen(":3000"))
}

func handleConnections(c *websocket.Conn) {
	defer func() {
		mu.Lock()
		delete(clients, c)
		mu.Unlock()
		c.Close()
	}()

	mu.Lock()
	clients[c] = true
	mu.Unlock()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Erro ao ler mensagem:", err)
			break
		}
		shared.Broadcast <- string(msg)
	}
}

func handleMessages() {
	for {
		msg := <-shared.Broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Printf("Erro ao enviar mensagem: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
