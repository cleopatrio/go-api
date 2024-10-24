package adapters

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/entities"
	"github.com/dock-tech/notes-api/internal/domain/usecases"
)

type NoteRepository interface {
	usecases.GetNoteUseCase
	usecases.ListNotesUseCase
	usecases.DeleteNoteUseCase

	Create(ctx context.Context, note entities.Note) (createdNote *entities.Note, err error)
}
