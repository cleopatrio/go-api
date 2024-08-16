package client

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	awsConfig map[string]*aws.Config
)

// getConfig creates aws configuration per region
func getConfig(region string) aws.Config {
	if awsConfig == nil {
		awsConfig = map[string]*aws.Config{}
	}
	if awsConfig[region] == nil {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			panic(err)
		}
		awsConfig[region] = &cfg
	}

	return *awsConfig[region]
}
