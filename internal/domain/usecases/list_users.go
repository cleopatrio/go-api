package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entities"
)

type ListUsersUseCase interface {
	List(ctx context.Context) (users []*entities.User, err error)
}
