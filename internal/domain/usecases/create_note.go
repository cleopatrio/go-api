package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entity"
)

type CreateNoteUseCase interface {
	Create(ctx context.Context, note entity.Note) (createdNote *entity.Note, err error)
}
