package servers

import (
	"github.com/dock-tech/notes-api/internal/delivery/adapters"
	"os"

	"github.com/gofiber/fiber/v3"
)

type server struct {
	usersController adapters.UsersController
	notesController adapters.NotesController
}

func (s server) Serve() {
	app := fiber.New()
	s.route(app)
	app.Listen(":" + os.Getenv("SERVER_PORT"))
}

func NewServer(usersController adapters.UsersController, notesController adapters.NotesController) adapters.Server {
	return &server{
		usersController: usersController,
		notesController: notesController,
	}
}
