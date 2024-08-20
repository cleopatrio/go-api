package interfaces

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/models"
)

type UseCase interface {
	ListUsers(ctx context.Context) (users *[]models.User, err error)
	GetUser(ctx context.Context, id string) (user *models.User, err error)
	CreateUser(ctx context.Context, user models.User) (createdUser *models.User, err error)
	DeleteUser(ctx context.Context, id string) (err error)

	ListNotes(ctx context.Context, userId string) (notes *[]models.Note, err error)
	GetNote(ctx context.Context, userId, noteId string) (note *models.Note, err error)
	CreateNote(ctx context.Context, note models.Note) (createdNote *models.Note, err error)
	DeleteNote(ctx context.Context, userId, noteId string) (err error)
}
