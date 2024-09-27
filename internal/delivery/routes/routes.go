package routes

import (
	"errors"

	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/gofiber/fiber/v3"
)

func responseParser(c fiber.Ctx, body any, statusCode int, err error) error {
	if err != nil {
		var errParsed *exceptions.ErrorType
		if !errors.As(err, errParsed) {
			errParsed = exceptions.NewInternalServerError(err.Error())
		}
		return c.Status(errParsed.StatusCode).JSON(errParsed.JSON())
	}
	return c.Status(statusCode).JSON(body)
}

type Router struct {
	controller interfaces.Controller
}

func (r Router) Route(app *fiber.App) {
	groupV1 := app.Group("/v1")

	groupV1.Get("/users", func(c fiber.Ctx) error {
		res, status, err := r.controller.ListUsers(c.Context())
		return responseParser(c, res, status, err)
	})

	groupV1.Get("/users/:id", func(c fiber.Ctx) error {
		userId := c.Params("id")
		res, status, err := r.controller.GetUser(c.Context(), userId)
		return responseParser(c, res, status, err)
	})

	groupV1.Post("/users", func(c fiber.Ctx) error {
		res, status, err := r.controller.CreateUser(c.Context(), c.Body())
		return responseParser(c, res, status, err)
	})

	groupV1.Delete("/users/:id", func(c fiber.Ctx) error {
		userId := c.Params("id")
		status, err := r.controller.DeleteUser(c.Context(), userId)
		return responseParser(c, nil, status, err)
	})

	groupV1.Get("/users/:id/notes", func(c fiber.Ctx) error {
		userId := c.Params("id")
		res, status, err := r.controller.ListNotes(c.Context(), userId)
		return responseParser(c, res, status, err)
	})

	groupV1.Get("/users/:id/notes/:note_id", func(c fiber.Ctx) error {
		userId := c.Params("id")
		noteId := c.Params("note_id")
		res, status, err := r.controller.GetNote(c.Context(), userId, noteId)
		return responseParser(c, res, status, err)
	})

	groupV1.Post("/users/:id/notes", func(c fiber.Ctx) error {
		userId := c.Params("id")
		res, status, err := r.controller.CreateNote(c.Context(), userId, c.Body())
		return responseParser(c, res, status, err)
	})

}

func NewRouter(controller interfaces.Controller) Router {
	return Router{controller: controller}
}
