package interfaces

import (
	"context"
)

type Controller interface {
	ListUsers(ctx context.Context) (response []byte, status int)
	GetUser(ctx context.Context, id string) (response []byte, status int)
	CreateUser(ctx context.Context, body []byte) (response []byte, status int)
	DeleteUser(ctx context.Context, id string) (response []byte, status int)

	ListNotes(ctx context.Context, userId string) (response []byte, status int)
	GetNote(ctx context.Context, userId, noteId string) (response []byte, status int)
	CreateNote(ctx context.Context, userId string, body []byte) (response []byte, status int)
	DeleteNote(ctx context.Context, noteId, userId string) (response []byte, status int)
}
