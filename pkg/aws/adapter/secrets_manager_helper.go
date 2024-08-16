package awsadapter

import "context"

type SecretsManagerHelper interface {
	GetSecret(ctx context.Context, secret *string, value any) error
}
