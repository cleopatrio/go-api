package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entity"
)

type ListNotesUseCase interface {
	List(ctx context.Context, userId string) (notes []*entity.Note, err error)
}
