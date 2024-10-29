package lambdago

import (
	"context"
	"jdvpl/awsgo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaExcuter)
}

func LambdaExcuter(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AwsStart()
	return event, nil
}
