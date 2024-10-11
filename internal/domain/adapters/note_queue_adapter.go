package adapters

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/entities"
)

type NoteQueue interface {
	Publish(ctx context.Context, note entities.Note) (err error)
}
