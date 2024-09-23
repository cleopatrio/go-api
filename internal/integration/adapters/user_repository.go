package adapters

import "github.com/dock-tech/notes-api/internal/domain/usecases"

type UserRepository interface {
	usecases.GetUserUseCase
	usecases.ListUsersUseCase
	usecases.CreateUserUseCase
	usecases.DeleteUserUseCase
}
