package lambdago

import (
	"context"
	"errors"
	"fmt"
	"jdvpl/awsgo"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaExcuter)
}

func LambdaExcuter(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AwsStart()
	// Get the secret from the AWS Secret Manager
	if !ParameterValidation() {
		fmt.Println("Missing SECRETNAME environment variable")
		err := errors.New("missing SECRETNAME environment variable")
		return event, err
	}
	return event, nil
}

func ParameterValidation() bool {
	var getParameter bool
	_, getParameter = os.LookupEnv("SECRETNAME")
	return getParameter
}
