package interfaces

import (
	"context"
)

type Controller interface {
	ListUsers(ctx context.Context) (response []byte, status int, err error)
	GetUser(ctx context.Context, id string) (response []byte, status int, err error)
	CreateUser(ctx context.Context, body []byte) (response []byte, status int, err error)
	DeleteUser(ctx context.Context, id string) (status int, err error)

	ListNotes(ctx context.Context, userId string) (response []byte, status int, err error)
	GetNote(ctx context.Context, userId, noteName string) (response []byte, status int, err error)
	CreateNote(ctx context.Context, userId string, noteName string) (response []byte, status int, err error)
	DeleteNote(ctx context.Context, userId string) (status int, err error)
}
