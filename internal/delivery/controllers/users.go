package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dock-tech/notes-api/internal/delivery/adapters"
	"github.com/dock-tech/notes-api/internal/delivery/validations"
	"github.com/dock-tech/notes-api/internal/domain/entity"
	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/domain/usecases"
	"log/slog"
	"net/http"
)

type userController struct {
	createUserUseCase   usecases.CreateUserUseCase
	deleteUserUseCase   usecases.DeleteUserUseCase
	getUserUseCase      usecases.GetUserUseCase
	listUsersUseCase    usecases.ListUsersUseCase
	errorHandlerUsecase adapters.ErrorHandler
}

func (u *userController) deferHandler(ctx context.Context, response *[]byte, status *int) {
	r := recover()
	if r != nil {
		*response, *status = u.errorHandlerUsecase.HandlePanic(ctx, r)
	}
}

func (u *userController) CreateUser(ctx context.Context, body []byte) (response []byte, status int) {
	defer u.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.CreateUser",
		slog.String("details", "process started"),
		slog.String("body", string(body)),
	)

	var user entity.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		err = exceptions.NewValidationError(fmt.Sprintf("error parsing json to user: %s", err.Error()))
		return u.errorHandlerUsecase.HandleError(ctx, err)
	}

	if err = validations.Validate(&user); err != nil {
		return u.errorHandlerUsecase.HandleError(ctx, err)
	}

	createdUser, err := u.createUserUseCase.Create(ctx, user)
	if err != nil {
		return u.errorHandlerUsecase.HandleError(ctx, err)
	}

	status = http.StatusCreated

	response, err = json.Marshal(createdUser)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing user to JSON: %s", err.Error()))
		return u.errorHandlerUsecase.HandleError(ctx, err)
	}

	slog.InfoContext(
		ctx, "controller.CreateNote",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func (u *userController) DeleteUser(ctx context.Context, id string) (response []byte, status int) {
	defer u.deferHandler(ctx, &response, &status)
	slog.InfoContext(ctx, "controller.DeleteNote",
		slog.String("details", "process started"),
		slog.String("userId", id),
	)

	err := u.deleteUserUseCase.Delete(ctx, id)
	if err != nil {
		return u.errorHandlerUsecase.HandleError(ctx, err)
	}

	status = http.StatusNoContent

	slog.InfoContext(
		ctx, "controller.DeleteNote",
		slog.String("details", "process finished"),
		slog.Int("status", status),
	)
	return
}

func (u *userController) GetUser(ctx context.Context, id string) (response []byte, status int) {
	defer u.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.GetUser",
		slog.String("details", "process started"),
		slog.String("userId", id),
	)

	user, err := u.getUserUseCase.Get(ctx, id)
	if err != nil {
		return u.errorHandlerUsecase.HandleError(ctx, err)
	}

	response, err = json.Marshal(user)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing user to JSON: %s", err.Error()))
		return u.errorHandlerUsecase.HandleError(ctx, err)
	}

	status = http.StatusOK

	slog.InfoContext(
		ctx, "controller.GetUser",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func (u *userController) ListUsers(ctx context.Context) (response []byte, status int) {
	defer u.deferHandler(ctx, &response, &status)

	slog.InfoContext(ctx, "controller.GetUser",
		slog.String("details", "process started"),
	)

	users, err := u.listUsersUseCase.List(ctx)
	if err != nil {
		return u.errorHandlerUsecase.HandleError(ctx, err)
	}

	response, err = json.Marshal(users)
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("error parsing users to JSON: %s", err.Error()))
		return u.errorHandlerUsecase.HandleError(ctx, err)
	}

	status = http.StatusOK

	slog.InfoContext(
		ctx, "controller.GetUser",
		slog.String("details", "process finished"),
		slog.String("response", string(response)),
		slog.Int("status", status),
	)
	return
}

func NewUsersController(
	createUserUseCase usecases.CreateUserUseCase,
	deleteUserUseCase usecases.DeleteUserUseCase,
	getUserUseCase usecases.GetUserUseCase,
	listUsersUseCase usecases.ListUsersUseCase,
	errorHandlerUsecase adapters.ErrorHandler,
) adapters.UsersController {
	return &userController{
		createUserUseCase:   createUserUseCase,
		deleteUserUseCase:   deleteUserUseCase,
		getUserUseCase:      getUserUseCase,
		listUsersUseCase:    listUsersUseCase,
		errorHandlerUsecase: errorHandlerUsecase,
	}
}
