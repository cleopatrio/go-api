package secrets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
)

type SecretClient interface {
	GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
}

type secret struct {
	client SecretClient
}

func (s secret) Get(ctx context.Context, key string) ([]byte, error) {
	res, err := s.client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return []byte(*res.SecretString), nil
}

func NewSecret(client SecretClient) interfaces.Secret {
	return &secret{client: client}
}
