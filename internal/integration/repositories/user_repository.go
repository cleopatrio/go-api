package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/dock-tech/notes-api/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type user struct {
	connection *gorm.DB
}

func (n user) Get(ctx context.Context, userId string) (user *models.User, err error) {
	err = n.connection.Where(
		&models.User{
			Id: userId,
		}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = exceptions.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId))
	}
	return
}

func (n user) Create(ctx context.Context, user models.User) (cratedUser *models.User, err error) {
	cratedUser = &user
	cratedUser.Id = uuid.NewString()
	now := time.Now()
	cratedUser.CreatedAt = &now
	cratedUser.UpdatedAt = &now
	err = n.connection.Create(&user).Error
	return
}

func (n user) Delete(ctx context.Context, userId string) (err error) {
	tx := n.connection.Delete(&models.User{Id: userId})
	err = tx.Error
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("failed to delete user with id %s", userId), err.Error())
	}

	if tx.RowsAffected == 0 {
		err = exceptions.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId))
	}

	return
}

func (n user) List(ctx context.Context) (users []*models.User, err error) {
	err = n.connection.Find(&users).Error
	return
}

func NewUser(connection *gorm.DB) interfaces.UserRepository {
	return &user{connection: connection}
}
