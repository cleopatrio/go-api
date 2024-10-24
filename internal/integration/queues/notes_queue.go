package queues

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/dock-tech/notes-api/internal/config/properties"
	"github.com/dock-tech/notes-api/internal/delivery/dtos"
	"github.com/dock-tech/notes-api/internal/domain/adapters"
	"github.com/dock-tech/notes-api/internal/domain/entities"
	"github.com/dock-tech/notes-api/internal/domain/exceptions"
)

type SqsClient interface {
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

type notes struct {
	sqsClient SqsClient
}

func (n *notes) Publish(ctx context.Context, note entities.Note) (err error) {
	var noteDTO dtos.Note

	bytes, err := json.Marshal(noteDTO.FromEntity(&note))
	if err != nil {
		return exceptions.NewNotesQueueError(err.Error())
	}

	_, err = n.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(bytes)),
		QueueUrl:    aws.String(properties.GetNotesQueueURL()),
	})

	if err != nil {
		return exceptions.NewNotesQueueError(err.Error())
	}

	return nil

}

func NewNotesQueue(sqsClient SqsClient) adapters.NoteQueue {
	return &notes{sqsClient}
}
