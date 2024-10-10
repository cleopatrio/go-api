package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entities"
)

type CreateUserUseCase interface {
	Create(ctx context.Context, user entities.User) (createdUser *entities.User, err error)
}
