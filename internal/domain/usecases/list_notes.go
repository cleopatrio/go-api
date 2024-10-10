package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entities"
)

type ListNotesUseCase interface {
	List(ctx context.Context, userId string) (notes []*entities.Note, err error)
}
