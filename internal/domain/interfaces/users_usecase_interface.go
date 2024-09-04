package interfaces

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/models"
)

type UsersUseCase interface {
	List(ctx context.Context) (users []*models.User, err error)
	Get(ctx context.Context, id string) (user *models.User, err error)
	Create(ctx context.Context, user models.User) (createdUser *models.User, err error)
	Delete(ctx context.Context, id string) (err error)
}
