package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/dock-tech/notes-api/internal/domain/entity"
	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/integration/adapters"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type userRepository struct {
	connection *gorm.DB
}

func (u *userRepository) Get(ctx context.Context, userId string) (user *entity.User, err error) {
	err = u.connection.WithContext(ctx).Where(
		&entity.User{
			Id: userId,
		}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = exceptions.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId))
	}
	return
}

func (u *userRepository) Create(ctx context.Context, user entity.User) (cratedUser *entity.User, err error) {
	cratedUser = &user
	cratedUser.Id = uuid.NewString()
	now := time.Now()
	cratedUser.CreatedAt = &now
	cratedUser.UpdatedAt = &now
	err = u.connection.WithContext(ctx).Create(&user).Error
	return
}

func (u *userRepository) Delete(ctx context.Context, userId string) (err error) {
	tx := u.connection.WithContext(ctx).Delete(&entity.User{Id: userId})
	err = tx.Error
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("failed to delete user with id %s", userId), err.Error())
	}

	if tx.RowsAffected == 0 {
		err = exceptions.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId))
	}

	return
}

func (u *userRepository) List(ctx context.Context) (users []*entity.User, err error) {
	err = u.connection.WithContext(ctx).Find(&users).Error
	return
}

func NewUser(connection *gorm.DB) adapters.UserRepository {
	return &userRepository{connection: connection}
}
