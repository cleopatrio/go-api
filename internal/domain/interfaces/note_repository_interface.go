package interfaces

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/models"
)

type NoteRepository interface {
	Get(ctx context.Context, userId string, noteId string) (*models.Note, error)
	Create(ctx context.Context, note models.Note) (createdNote *models.Note, err error)
	Delete(ctx context.Context, noteId string) error
	List(ctx context.Context, userId string) (notes []*models.Note, err error)
}
