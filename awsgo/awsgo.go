package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// Ctx is the context for the AWS SDK
// Cfg is the configuration for the AWS SDK
// err is the error for the AWS SDK
var Ctx context.Context
var Cfg aws.Config
var err error

func AwsStart() {
	// Create a context to use in the service's requests with TODO
	Ctx = context.TODO()
	// Load the SDK's configuration from the environment and shared config, and create a new service client
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion(("us-east-1")))
	if err != nil {
		panic("configuration error, " + err.Error())
	}
}
