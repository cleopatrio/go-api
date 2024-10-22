package adapters

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/entities"
)

type CreateNoteRepository interface {
	Create(ctx context.Context, note entities.Note) (createdNote *entities.Note, err error)
}
