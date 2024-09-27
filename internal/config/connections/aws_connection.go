package connections

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/dock-tech/notes-api/internal/config/properties"
	"github.com/dock-tech/notes-api/internal/integration/secrets"
	"os"
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
