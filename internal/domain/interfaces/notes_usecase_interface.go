package interfaces

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/models"
)

type NotesUseCase interface {
	List(ctx context.Context, userId string) (notes []*models.Note, err error)
	Get(ctx context.Context, userId, noteId string) (note *models.Note, err error)
	Create(ctx context.Context, note models.Note) (createdNote *models.Note, err error)
	Delete(ctx context.Context, userId, noteId string) (err error)
}
