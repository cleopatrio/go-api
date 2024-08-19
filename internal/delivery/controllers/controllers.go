package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dock-tech/notes-api/internal/delivery/validations"
	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/dock-tech/notes-api/internal/domain/models"
)

type controller struct {
	usecase interfaces.UseCase
}

func (c *controller) CreateNote(ctx context.Context, userId string, body []byte) (response []byte, status int, err error) {
	slog.InfoContext(ctx, "controller.CreateNote",
		slog.String("details", "process started"),
		slog.String("userId", userId),
	)

	var note models.Note
	err = json.Unmarshal(body, &note)
	if err != nil {
		err = exceptions.NewValidationError(fmt.Sprintf("error parsing JSON to note: %s", err.Error()))
		return
	}

	if err = validations.Validate(&note); err != nil {
		return
	}

	createdNote, err := c.usecase.CreateNote(ctx, note)
	if err != nil {
		return
	}

	if response, err = json.Marshal(createdNote); err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing note to JSON: %s", err.Error()))
		return
	}

	status = http.StatusCreated

	slog.InfoContext(
		ctx, "controller.CreateNote",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func (c *controller) CreateUser(ctx context.Context, body []byte) (response []byte, status int, err error) {
	slog.InfoContext(ctx, "controller.CreateUser",
		slog.String("details", "process started"),
		slog.String("body", string(body)),
	)

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		err = exceptions.NewValidationError(fmt.Sprintf("error parsing json to user: %s", err.Error()))
		return
	}

	createdUser, err := c.usecase.CreateUser(ctx, user)
	if err != nil {
		return
	}

	status = http.StatusCreated

	response, err = json.Marshal(createdUser)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing user to JSON: %s", err.Error()))
		return
	}

	slog.InfoContext(
		ctx, "controller.CreateNote",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func (c *controller) DeleteNote(ctx context.Context, userId, noteId string) (status int, err error) {
	slog.InfoContext(ctx, "controller.DeleteNote",
		slog.String("details", "process started"),
		slog.String("noteId", noteId),
		slog.String("userId", userId),
	)

	err = c.usecase.DeleteNote(ctx, userId, noteId)
	if err != nil {
		return
	}

	status = http.StatusNoContent

	slog.InfoContext(
		ctx, "controller.DeleteNote",
		slog.String("details", "process finished"),
		slog.Int("status", status),
	)
	return
}

func (c *controller) DeleteUser(ctx context.Context, id string) (status int, err error) {
	slog.InfoContext(ctx, "controller.DeleteNote",
		slog.String("details", "process started"),
		slog.String("userId", id),
	)

	err = c.usecase.DeleteUser(ctx, id)
	if err != nil {
		return
	}

	status = http.StatusNoContent

	slog.InfoContext(
		ctx, "controller.DeleteNote",
		slog.String("details", "process finished"),
		slog.Int("status", status),
	)
	return
}

func (c *controller) GetNote(ctx context.Context, userId string, noteId string) (response []byte, status int, err error) {
	slog.InfoContext(ctx, "controller.GetNote",
		slog.String("details", "process started"),
		slog.String("userId", userId),
	)

	note, err := c.usecase.GetNote(ctx, userId, noteId)
	if err != nil {
		return
	}

	response, err = json.Marshal(note)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing note to JSON: %s", err.Error()))
		return
	}

	status = http.StatusOK

	slog.InfoContext(
		ctx, "controller.GetNote",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func (c *controller) GetUser(ctx context.Context, id string) (response []byte, status int, err error) {
	slog.InfoContext(ctx, "controller.GetUser",
		slog.String("details", "process started"),
		slog.String("userId", id),
	)

	user, err := c.usecase.GetUser(ctx, id)
	if err != nil {
		return
	}

	response, err = json.Marshal(user)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing user to JSON: %s", err.Error()))
		return
	}

	status = http.StatusOK

	slog.InfoContext(
		ctx, "controller.GetUser",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func (c *controller) ListNotes(ctx context.Context, userId string) (response []byte, status int, err error) {
	slog.InfoContext(ctx, "controller.GetUser",
		slog.String("details", "process started"),
		slog.String("userId", userId),
	)

	notes, err := c.usecase.ListNotes(ctx, userId)
	if err != nil {
		return
	}

	response, err = json.Marshal(notes)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing notes to JSON: %s", err.Error()))
		return
	}

	status = http.StatusOK

	slog.InfoContext(
		ctx, "controller.GetUser",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func (c *controller) ListUsers(ctx context.Context) (response []byte, status int, err error) {
	slog.InfoContext(ctx, "controller.GetUser",
		slog.String("details", "process started"),
	)

	users, err := c.usecase.ListUsers(ctx)
	if err != nil {
		return
	}

	response, err = json.Marshal(users)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing users to JSON: %s", err.Error()))
		return
	}

	status = http.StatusOK

	slog.InfoContext(
		ctx, "controller.GetUser",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func NewController(usecase interfaces.UseCase) interfaces.Controller {
	return &controller{
		usecase: usecase,
	}
}
