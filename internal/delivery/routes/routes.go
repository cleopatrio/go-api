package routes

import (
	"errors"

	"github.com/dock-tech/notes-api/internal/domain/exceptions"
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

func Route(app *fiber.App) {
	groupV1 := app.Group("/v1")

	groupV1.Get("/users", func(c fiber.Ctx) error {
		users := []string{"John", "Jane", "Bob"}
		return c.JSON(users)
	})

	groupV1.Get("/users/:id", func(c fiber.Ctx) error {
		// Get user by ID
		userID := c.Params("id")
		// Your code to fetch user from database or any other data source

		// Return user as JSON response
		return c.JSON(user)
	})

	groupV1.Post("/users", func(c fiber.Ctx) error {
		// Parse request body to get user data
		var newUser User
		if err := c.BodyParser(&newUser); err != nil {
			return err
		}
		// Your code to create a new user in the database or any other data source

		// Return created user as JSON response
		return c.JSON(newUser)
	})

	groupV1.Delete("/users/:id", func(c fiber.Ctx) error {
		// Get user ID from URL parameter
		userID := c.Params("id")
		// Your code to fetch user from database or any other data source

		// Your code to delete the user from the database or any other data source

		// Return success message as JSON response
		return c.JSON(fiber.Map{"message": "User deleted successfully"})
	})

	groupV1.Get("/users/:id/notes", func(c fiber.Ctx) error {
		// Get user ID from URL parameter
		userID := c.Params("id")
		// Your code to fetch user's notes from the database or any other data source

		// Return user's notes as JSON response
		return c.JSON(notes)
	})

	groupV1.Get("/users/:id/notes/:note_id", func(c fiber.Ctx) error {
		// Get user ID from URL parameter
		userID := c.Params("id")
		// Your code to fetch user's notes from the database or any other data source

		// Return user's notes as JSON response
		return c.JSON(notes)
	})
	groupV1.Post("/users/:id/notes", func(c fiber.Ctx) error {
		// Get user ID from URL parameter
		userID := c.Params("id")
		// Parse request body to get note data
		var newNote Note
		if err := c.BodyParser(&newNote); err != nil {
			return err
		}
		// Your code to create a new note for the user in the database or any other data source

		// Return created note as JSON response
		return c.JSON(newNote)
	})

}
