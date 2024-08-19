package usecases

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/dock-tech/notes-api/internal/domain/models"
)

type usecase struct {
}


func (u *usecase) CreateNote(ctx context.Context, note models.Note) (createdNote *models.Note, err error) {
	panic("unimplemented")
}

// CreateUser implements interfaces.UseCase.
func (u *usecase) CreateUser(ctx context.Context, user models.User) (createdUser *models.User, err error) {
	panic("unimplemented")
}

// DeleteNote implements interfaces.UseCase.
func (u *usecase) DeleteNote(ctx context.Context, userId string, noteId string) (err error) {
	panic("unimplemented")
}

// DeleteUser implements interfaces.UseCase.
func (u *usecase) DeleteUser(ctx context.Context, id string) (err error) {
	panic("unimplemented")
}

// GetNote implements interfaces.UseCase.
func (u *usecase) GetNote(ctx context.Context, userId string, noteId string) (note *models.Note, err error) {
	panic("unimplemented")
}

// GetUser implements interfaces.UseCase.
func (u *usecase) GetUser(ctx context.Context, id string) (user *models.User, err error) {
	panic("unimplemented")
}

// ListNotes implements interfaces.UseCase.
func (u *usecase) ListNotes(ctx context.Context, userId string) (notes *[]models.Note, err error) {
	panic("unimplemented")
}

// ListUsers implements interfaces.UseCase.
func (u *usecase) ListUsers(ctx context.Context) (users *[]models.User, err error) {
	panic("unimplemented")
}

func NewUsecase() interfaces.UseCase {
	return &usecase{}
}
