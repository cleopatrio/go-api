package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/dock-tech/notes-api/internal/domain/entity"
	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/integration/adapters"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	connection *gorm.DB
}

func (u *userRepository) Get(ctx context.Context, userId string) (*entity.User, error) {
	var user *entity.User
	err := u.connection.WithContext(ctx).Where(
		&entity.User{
			Id: userId,
		}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exceptions.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId))
	}

	return user, nil
}

func (u *userRepository) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	err := u.connection.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, exceptions.NewInternalServerError("failed to create user", err.Error())
	}

	return &user, nil
}

func (u *userRepository) Delete(ctx context.Context, userId string) error {
	tx := u.connection.WithContext(ctx).Select(clause.Associations).Delete(&entity.User{Id: userId})
	err := tx.Error
	if err != nil {
		return exceptions.NewInternalServerError(fmt.Sprintf("failed to delete user with id %s", userId), err.Error())
	}

	if tx.RowsAffected == 0 {
		err = exceptions.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId))
	}

	return nil
}

func (u *userRepository) List(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	err := u.connection.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, exceptions.NewInternalServerError("failed to list users", err.Error())
	}

	return users, nil
}

func NewUser(connection *gorm.DB) adapters.UserRepository {
	return &userRepository{connection: connection}
}
