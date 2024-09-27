package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entity"
)

type GetUserUseCase interface {
	Get(ctx context.Context, id string) (user *entity.User, err error)
}
