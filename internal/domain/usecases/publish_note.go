package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entity"
)

// TODO implementar a interface
type PublishNoteUseCase interface {
	Publish(ctx context.Context, note entity.Note) (err error)
}
