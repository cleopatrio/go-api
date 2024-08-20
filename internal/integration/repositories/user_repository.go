package repositories

import (
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/dock-tech/notes-api/internal/domain/models"
	"gorm.io/gorm"
)

type user struct {
	connection *gorm.DB
}

func (n user) Get(userId string) (user *models.User, err error) {
	err = n.connection.Where(
		&models.User{
			Id:     userId,
		}).First(&user).Error
	return
}

func (n user) Create(user models.User) (err error) {
	err = n.connection.Create(&user).Error
	return
}

func (n user) Delete(userId string) (err error) {
	err = n.connection.Delete(&models.User{}, userId).Error
	return
}

func (n user) List(userId string) (users []*models.User, err error) {
	err = n.connection.Where(&models.User{Id:userId}).Find(&users).Error
	return
}

func NewUser(connection *gorm.DB) interfaces.UserRepository {
	return &user{connection: connection}
}
