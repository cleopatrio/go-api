package awsadapter

import "context"

type SqsHelper interface {
	Send(ctx context.Context, message any, queue string) error
}
