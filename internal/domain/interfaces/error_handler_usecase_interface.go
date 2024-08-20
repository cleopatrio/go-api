package interfaces

import "context"

type ErrorHandlerUsecase interface {
	HandlePanic(ctx context.Context, recovered any) (res []byte, statusCode int)
	HandleError(ctx context.Context, err error) (res []byte, statusCode int)
}
