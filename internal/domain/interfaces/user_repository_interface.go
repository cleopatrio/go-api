package interfaces

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/models"
)

type UserRepository interface {
	Get(ctx context.Context, userId string) (user *models.User, err error)
	Create(ctx context.Context, user models.User) (createdUser *models.User, err error)
	Delete(ctx context.Context, userId string) error
	List(ctx context.Context) (users []*models.User, err error)
}
