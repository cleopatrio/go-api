package dtos

import (
	"time"

	"github.com/dock-tech/notes-api/internal/domain/entities"
)

type User struct {
	Id        string     `json:"id"`
	Name      string     `json:"name" validate:"required,min=3"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u User) ToEntity() entities.User {
	return entities.User{
		Id:        u.Id,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u User) FromEntity(user *entities.User) User {
	return User{
		Id:        user.Id,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type Users []*User

func (u Users) FromEntities(users []*entities.User) Users {
	u = make([]*User, len(users))
	for i, user := range users {
		u[i] = &User{
			Id:        user.Id,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}
	return u
}
