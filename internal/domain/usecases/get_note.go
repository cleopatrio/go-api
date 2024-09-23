package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entity"
)

type GetNoteUseCase interface {
	Get(ctx context.Context, userId, noteId string) (note *entity.Note, err error)
}
