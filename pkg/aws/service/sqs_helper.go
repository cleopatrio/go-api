package service

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	awsadapter "github.com/dock-tech/notes-api/pkg/aws/adapter"
)

type sqsHelper struct {
	sqs awsadapter.Sqs
}

func SqsHelper(sqs awsadapter.Sqs) awsadapter.SqsHelper {
	return &sqsHelper{
		sqs: sqs,
	}
}

func (s *sqsHelper) Send(ctx context.Context, message any, queue string) error {
	println("sqsHelper.Send Started")

	stringMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	payload := string(stringMessage)
	messageInput := &sqs.SendMessageInput{
		MessageBody: &payload,
		QueueUrl:    &queue,
	}

	if correlationId, ok := ctx.Value("correlationId").(string); ok && correlationId != "" {
		messageInput.MessageAttributes = map[string]types.MessageAttributeValue{
			"correlationId": {
				DataType:    aws.String("String"),
				StringValue: aws.String(correlationId),
			},
		}
	}

	println("sqsHelper.Send message: ", messageInput)

	result, err := s.sqs.SendMessage(ctx, messageInput)
	if err != nil {
		println("sqsHelper.Send Error: ", err.Error())
		return err
	}

	println("sqsHelper.Send Finished")
	println("sqsHelper.Send result: ", result)
	return nil
}
