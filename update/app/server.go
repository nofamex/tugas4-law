package app

import (
	"log"
	"net/http"

	"read/db"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	db *gorm.DB
}

func NewServer(db *gorm.DB) (*Server, error) {
	app := fiber.New()
	server := &Server{
		app: app,
		db: db,
	}

	log.Println("create server")
	server.setupRoutes()
	return server, nil
}

func (server *Server) setupRoutes() {
	server.app.Post("/update", func(c *fiber.Ctx) error {
		var requestBody UserModel
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": "please send the correct body"})
		}

		user := &db.User{
			Name: requestBody.Name,
			NPM: requestBody.NPM,
		}

		server.db.Create(user)

		return c.Status(http.StatusOK).JSON(fiber.Map{"status": "OK"})
	})
}
	

func (server *Server) StartServer(port string) error {
	log.Printf("server running on port %s", port)
	return server.app.Listen(":" + port)
}