package interfaces

import "github.com/dock-tech/notes-api/internal/domain/models"

type UserRepository interface {
	Get(userId string) (*models.User, error)
	Create(user models.User) error
	Delete(userId string) error
	List(userId string) ([]*models.User, error)
}
