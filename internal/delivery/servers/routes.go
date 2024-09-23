package servers

import (
	"github.com/gofiber/fiber/v3"
)

func respond(c fiber.Ctx, status int, res []byte) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(status).Send(res)
}

func (r server) route(app *fiber.App) {
	groupV1 := app.Group("/v1")

	groupV1.Get("/users", func(c fiber.Ctx) error {
		res, status := r.usersController.ListUsers(c.Context())
		return respond(c, status, res)
	})

	groupV1.Get("/users/:id", func(c fiber.Ctx) error {
		userId := c.Params("id")
		res, status := r.usersController.GetUser(c.Context(), userId)
		return respond(c, status, res)
	})

	groupV1.Post("/users", func(c fiber.Ctx) error {
		res, status := r.usersController.CreateUser(c.Context(), c.Body())
		return respond(c, status, res)
	})

	groupV1.Delete("/users/:id", func(c fiber.Ctx) error {
		userId := c.Params("id")
		res, status := r.usersController.DeleteUser(c.Context(), userId)
		return respond(c, status, res)
	})

	groupV1.Delete("/users/:id/notes/:note_id", func(c fiber.Ctx) error {
		userId := c.Params("id")
		noteId := c.Params("note_id")
		res, status := r.notesController.DeleteNote(c.Context(), noteId, userId)
		return respond(c, status, res)
	})
	groupV1.Get("/users/:id/notes", func(c fiber.Ctx) error {
		userId := c.Params("id")
		res, status := r.notesController.ListNotes(c.Context(), userId)
		return respond(c, status, res)
	})

	groupV1.Get("/users/:id/notes/:note_id", func(c fiber.Ctx) error {
		userId := c.Params("id")
		noteId := c.Params("note_id")
		res, status := r.notesController.GetNote(c.Context(), userId, noteId)
		return respond(c, status, res)
	})

	groupV1.Post("/users/:id/notes", func(c fiber.Ctx) error {
		userId := c.Params("id")
		res, status := r.notesController.CreateNote(c.Context(), userId, c.Body())
		return respond(c, status, res)
	})

}
