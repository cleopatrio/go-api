package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dock-tech/notes-api/internal/delivery/adapters"
	"github.com/dock-tech/notes-api/internal/delivery/validations"
	"github.com/dock-tech/notes-api/internal/domain/entity"
	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/domain/usecases"
	"log/slog"
	"net/http"
)

type notesController struct {
	createNoteUseCase   usecases.CreateNoteUseCase
	deleteNoteUseCase   usecases.DeleteNoteUseCase
	getNoteUseCase      usecases.GetNoteUseCase
	listNotesUseCase    usecases.ListNotesUseCase
	errorHandlerUsecase adapters.ErrorHandler
}

func (n *notesController) deferHandler(ctx context.Context, response *[]byte, status *int) {
	r := recover()
	if r != nil {
		*response, *status = n.errorHandlerUsecase.HandlePanic(ctx, r)
	}
}

func (n *notesController) CreateNote(ctx context.Context, userId string, body []byte) (response []byte, status int) {
	defer n.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.CreateNote",
		slog.String("details", "process started"),
		slog.String("userId", userId),
	)

	var note entity.Note
	err := json.Unmarshal(body, &note)
	if err != nil {
		err = exceptions.NewValidationError(fmt.Sprintf("error parsing JSON to note: %s", err.Error()))
		return n.errorHandlerUsecase.HandleError(ctx, err)
	}

	note.UserId = userId

	if err = validations.Validate(&note); err != nil {
		return n.errorHandlerUsecase.HandleError(ctx, err)
	}

	createdNote, err := n.createNoteUseCase.Create(ctx, note)
	if err != nil {
		return n.errorHandlerUsecase.HandleError(ctx, err)
	}

	if response, err = json.Marshal(createdNote); err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing note to JSON: %s", err.Error()))
		return n.errorHandlerUsecase.HandleError(ctx, err)
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

func (n *notesController) DeleteNote(ctx context.Context, noteId, userId string) (response []byte, status int) {
	defer n.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.DeleteNote",
		slog.String("details", "process started"),
		slog.String("noteId", noteId),
		slog.String("userId", userId),
	)

	err := n.deleteNoteUseCase.Delete(ctx, userId, noteId)
	if err != nil {
		return n.errorHandlerUsecase.HandleError(ctx, err)
	}

	status = http.StatusNoContent

	slog.InfoContext(
		ctx, "controller.DeleteNote",
		slog.String("details", "process finished"),
		slog.Int("status", status),
	)
	return
}

func (n *notesController) GetNote(ctx context.Context, userId string, noteId string) (response []byte, status int) {
	defer n.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.GetNote",
		slog.String("details", "process started"),
		slog.String("userId", userId),
	)

	note, err := n.getNoteUseCase.Get(ctx, userId, noteId)
	if err != nil {
		return n.errorHandlerUsecase.HandleError(ctx, err)
	}

	response, err = json.Marshal(note)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing note to JSON: %s", err.Error()))
		return n.errorHandlerUsecase.HandleError(ctx, err)
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

func (n *notesController) ListNotes(ctx context.Context, userId string) (response []byte, status int) {
	defer n.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.GetUser",
		slog.String("details", "process started"),
		slog.String("userId", userId),
	)

	notes, err := n.listNotesUseCase.List(ctx, userId)
	if err != nil {
		return n.errorHandlerUsecase.HandleError(ctx, err)
	}

	response, err = json.Marshal(notes)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing notes to JSON: %s", err.Error()))
		return n.errorHandlerUsecase.HandleError(ctx, err)
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

func NewNotesController(
	createNoteUseCase usecases.CreateNoteUseCase,
	deleteNoteUseCase usecases.DeleteNoteUseCase,
	getNoteUseCase usecases.GetNoteUseCase,
	listNoteUseCase usecases.ListNotesUseCase,
	errorHandlerUsecase adapters.ErrorHandler,
) adapters.NotesController {
	return &notesController{
		createNoteUseCase:   createNoteUseCase,
		deleteNoteUseCase:   deleteNoteUseCase,
		getNoteUseCase:      getNoteUseCase,
		listNotesUseCase:    listNoteUseCase,
		errorHandlerUsecase: errorHandlerUsecase,
	}
}
