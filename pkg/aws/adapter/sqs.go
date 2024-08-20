package awsadapter

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Sqs interface {
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}
