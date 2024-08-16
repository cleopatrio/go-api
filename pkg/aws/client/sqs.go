package client

import (
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"os"
)

func Sqs(region string) *sqs.Client {
	cfg := getConfig(region)
	if awsUrl := os.Getenv("AWS_URL"); awsUrl != "" {
		return sqs.NewFromConfig(cfg, func(o *sqs.Options) {
			o.BaseEndpoint = &awsUrl
		})
	}

	return sqs.NewFromConfig(cfg)
}
