package servers

import (
	"os"

	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/gofiber/fiber/v3"
)

type server struct {
	controller interfaces.Controller
}

func (s server) Serve() {
	app := fiber.New()
	s.route(app)
	app.Listen(":" + os.Getenv("PORT"))
}

func NewServer(controllers interfaces.Controller) interfaces.Server {
	return &server{controller: controllers}
}
