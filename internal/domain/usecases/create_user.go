package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entity"
)

type CreateUserUseCase interface {
	Create(ctx context.Context, user entity.User) (createdUser *entity.User, err error)
}
