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
	usersUseCase        interfaces.UsersUseCase
	notesUseCase        interfaces.NotesUseCase
	errorHandlerUsecase interfaces.ErrorHandlerUsecase
}

func (c *controller) deferHandler(ctx context.Context, response *[]byte, status *int) {
	r := recover()
	if r != nil {
		*response, *status = c.errorHandlerUsecase.HandlePanic(ctx, r)
	}
}

func (c *controller) CreateNote(ctx context.Context, userId string, body []byte) (response []byte, status int) {
	defer c.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.CreateNote",
		slog.String("details", "process started"),
		slog.String("userId", userId),
	)

	var note models.Note
	err := json.Unmarshal(body, &note)
	if err != nil {
		err = exceptions.NewValidationError(fmt.Sprintf("error parsing JSON to note: %s", err.Error()))
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	note.UserId = userId

	if err = validations.Validate(&note); err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	createdNote, err := c.notesUseCase.Create(ctx, note)
	if err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	if response, err = json.Marshal(createdNote); err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing note to JSON: %s", err.Error()))
		return c.errorHandlerUsecase.HandleError(ctx, err)
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

func (c *controller) CreateUser(ctx context.Context, body []byte) (response []byte, status int) {
	defer c.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.CreateUser",
		slog.String("details", "process started"),
		slog.String("body", string(body)),
	)

	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		err = exceptions.NewValidationError(fmt.Sprintf("error parsing json to user: %s", err.Error()))
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	if err = validations.Validate(&user); err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	createdUser, err := c.usersUseCase.Create(ctx, user)
	if err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	status = http.StatusCreated

	response, err = json.Marshal(createdUser)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing user to JSON: %s", err.Error()))
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	slog.InfoContext(
		ctx, "controller.CreateNote",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func (c *controller) DeleteNote(ctx context.Context, noteId, userId string) (response []byte, status int) {
	defer c.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.DeleteNote",
		slog.String("details", "process started"),
		slog.String("noteId", noteId),
		slog.String("userId", userId),
	)

	err := c.notesUseCase.Delete(ctx, userId, noteId)
	if err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	status = http.StatusNoContent

	slog.InfoContext(
		ctx, "controller.DeleteNote",
		slog.String("details", "process finished"),
		slog.Int("status", status),
	)
	return
}

func (c *controller) DeleteUser(ctx context.Context, id string) (response []byte, status int) {
	defer c.deferHandler(ctx, &response, &status)
	slog.InfoContext(ctx, "controller.DeleteNote",
		slog.String("details", "process started"),
		slog.String("userId", id),
	)

	err := c.usersUseCase.Delete(ctx, id)
	if err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	status = http.StatusNoContent

	slog.InfoContext(
		ctx, "controller.DeleteNote",
		slog.String("details", "process finished"),
		slog.Int("status", status),
	)
	return
}

func (c *controller) GetNote(ctx context.Context, userId string, noteId string) (response []byte, status int) {
	defer c.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.GetNote",
		slog.String("details", "process started"),
		slog.String("userId", userId),
	)

	note, err := c.notesUseCase.Get(ctx, userId, noteId)
	if err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	response, err = json.Marshal(note)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing note to JSON: %s", err.Error()))
		return c.errorHandlerUsecase.HandleError(ctx, err)
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

func (c *controller) GetUser(ctx context.Context, id string) (response []byte, status int) {
	defer c.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.GetUser",
		slog.String("details", "process started"),
		slog.String("userId", id),
	)

	user, err := c.usersUseCase.Get(ctx, id)
	if err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	response, err = json.Marshal(user)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing user to JSON: %s", err.Error()))
		return c.errorHandlerUsecase.HandleError(ctx, err)
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

func (c *controller) ListNotes(ctx context.Context, userId string) (response []byte, status int) {
	defer c.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.GetUser",
		slog.String("details", "process started"),
		slog.String("userId", userId),
	)

	notes, err := c.notesUseCase.List(ctx, userId)
	if err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	response, err = json.Marshal(notes)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing notes to JSON: %s", err.Error()))
		return c.errorHandlerUsecase.HandleError(ctx, err)
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

func (c *controller) ListUsers(ctx context.Context) (response []byte, status int) {
	defer c.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.GetUser",
		slog.String("details", "process started"),
	)

	users, err := c.usersUseCase.List(ctx)
	if err != nil {
		return c.errorHandlerUsecase.HandleError(ctx, err)
	}

	response, err = json.Marshal(users)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing users to JSON: %s", err.Error()))
		return c.errorHandlerUsecase.HandleError(ctx, err)
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

func NewController(notesUseCase interfaces.NotesUseCase, usersUseCase interfaces.UsersUseCase, errorHandlerUsecase interfaces.ErrorHandlerUsecase) interfaces.Controller {
	return &controller{
		errorHandlerUsecase: errorHandlerUsecase,
		usersUseCase:        usersUseCase,
		notesUseCase:        notesUseCase,
	}
}
