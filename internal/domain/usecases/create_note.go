package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entities"
)

type CreateNoteUseCase interface {
	Create(ctx context.Context, note entities.Note) (createdNote *entities.Note, err error)
}
