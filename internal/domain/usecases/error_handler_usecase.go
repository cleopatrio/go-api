package usecases

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
)

type errorHandler struct {
}

func (e *errorHandler) HandlePanic(ctx context.Context, recovered any) (response []byte, statusCode int) {
	if recovered != nil {
		err := exceptions.NewInternalServerError(fmt.Sprintf("panic: %v", recovered))
		return e.HandleError(ctx, err)
	}
	return
}

func (e *errorHandler) HandleError(ctx context.Context, err error) (response []byte, statusCode int) {
	errParsed := &exceptions.ErrorType{}
	if !errors.As(err, errParsed) {
		errParsed = exceptions.NewInternalServerError(err.Error())
	}

	slog.ErrorContext(ctx, "errorHandler.HandleError", slog.String("errorDetails", string(errParsed.JSON())))

	/*
	   if errParsed.StatusCode == http.StatusInternalServerError {
	       slack notification?
	   }
	*/

	return errParsed.JSON(), errParsed.StatusCode
}

func NewErrorHandler() interfaces.ErrorHandlerUsecase {
	return &errorHandler{}
}
