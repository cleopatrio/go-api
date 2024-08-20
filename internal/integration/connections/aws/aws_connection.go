package connections

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"os"
)

func GetAwsConfig() aws.Config {
	var awsConfig *aws.Config
	if os.Getenv("AWS_URL") != "" {
		newAwsConfig, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithEndpointResolverWithOptions(newAwsEndpointResolver()),
		)
		if err != nil {
			panic("configuration error, " + err.Error())
		}
		awsConfig = &newAwsConfig
	} else {
		newAwsConfig, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("configuration error, " + err.Error())
		}
		awsConfig = &newAwsConfig
	}

	return *awsConfig
}

// newEndpointResolver creates a new aws endpoint. Can override the endpoint when used with localstack.
func newAwsEndpointResolver() aws.EndpointResolverWithOptionsFunc {
	return func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:       "aws",
			URL:               os.Getenv("AWS_URL"),
			SigningRegion:     os.Getenv("AWS_REGION"),
			HostnameImmutable: true,
		}, nil
	}
}
