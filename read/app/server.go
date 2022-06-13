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
	server.app.Get("/read/:npm", func(c *fiber.Ctx) error {
		var npm string = c.Params("npm")

		var user db.User

		server.db.Where("npm = ?", npm).Find(&user)

		return c.Status(http.StatusOK).JSON(fiber.Map{"status": "OK", "npm": user.NPM, "name": user.Name})
	})
}
	

func (server *Server) StartServer(port string) error {
	log.Printf("server running on port %s", port)
	return server.app.Listen(":" + port)
}