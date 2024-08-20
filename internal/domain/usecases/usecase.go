package usecases

import (
	"context"

	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/dock-tech/notes-api/internal/domain/models"
)

type usecase struct {
	noteRepository interfaces.NoteRepository
	userRepository interfaces.UserRepository
}

func (u *usecase) CreateNote(ctx context.Context, note models.Note) (createdNote *models.Note, err error) {
	createdNote, err = u.noteRepository.Create(ctx, note)
	return
}

func (u *usecase) CreateUser(ctx context.Context, user models.User) (createdUser *models.User, err error) {
	createdUser, err = u.userRepository.Create(ctx, user)
	return
}

func (u *usecase) DeleteNote(ctx context.Context, userId string, noteId string) (err error) {
	err = u.noteRepository.Delete(ctx, noteId)
	return
}

func (u *usecase) DeleteUser(ctx context.Context, id string) (err error) {
	err = u.userRepository.Delete(ctx, id)
	return
}

func (u *usecase) GetNote(ctx context.Context, userId string, noteId string) (note *models.Note, err error) {
	note, err = u.noteRepository.Get(ctx, userId, noteId)
	return
}

func (u *usecase) GetUser(ctx context.Context, id string) (user *models.User, err error) {
	user, err = u.userRepository.Get(ctx, id)
	return
}

func (u *usecase) ListNotes(ctx context.Context, userId string) (notes []*models.Note, err error) {
	notes, err = u.noteRepository.List(ctx, userId)
	return
}

func (u *usecase) ListUsers(ctx context.Context) (users []*models.User, err error) {
	users, err = u.userRepository.List(ctx)
	return
}

func NewUsecase(noteRepository interfaces.NoteRepository, userRepository interfaces.UserRepository) interfaces.UseCase {
	return &usecase{
		noteRepository: noteRepository,
		userRepository: userRepository,
	}
}
