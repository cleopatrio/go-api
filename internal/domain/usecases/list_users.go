package usecases

import (
	"context"
	"github.com/dock-tech/notes-api/internal/domain/entity"
)

type ListUsersUseCase interface {
	List(ctx context.Context) (users []*entity.User, err error)
}
