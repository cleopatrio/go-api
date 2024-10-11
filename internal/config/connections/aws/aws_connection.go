package aws

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/dock-tech/notes-api/internal/config/properties"
	"github.com/dock-tech/notes-api/internal/integration/secrets"
)

func NewAws() aws.Config {
	awsConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(properties.GetRegion()))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	return awsConfig
}

func NewAwsSecretsManager(cfg aws.Config) secrets.SecretClient {
	if awsUrl := os.Getenv("AWS_URL"); awsUrl != "" {
		println("Using AWS_URL: ", awsUrl)
		return secretsmanager.NewFromConfig(cfg, func(o *secretsmanager.Options) {
			o.BaseEndpoint = &awsUrl
		})
	}

	return secretsmanager.NewFromConfig(cfg)
}

func NewAwsSqs(cfg aws.Config) *sqs.Client {
	if awsUrl := os.Getenv("AWS_URL"); awsUrl != "" {
		println("Using AWS_URL: ", awsUrl)
		return sqs.NewFromConfig(cfg, func(o *sqs.Options) {
			o.BaseEndpoint = &awsUrl
		})
	}

	return sqs.NewFromConfig(cfg)
}
