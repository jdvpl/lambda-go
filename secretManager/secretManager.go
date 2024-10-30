package secretManager

import (
	"encoding/json"
	"fmt"
	"jdvpl/awsgo"
	"jdvpl/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// GetSecret retrieves a secret from AWS Secrets Manager and unmarshals it into a models.SecretRDSJson struct.
func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	fmt.Println("Getting secret from AWS Secret Manager: " + secretName)

	// Create a new Secrets Manager client using the configuration from awsgo package.
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	// Get the secret value from AWS Secrets Manager.
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName), // & is used to get the address of the secretName string.
	})
	if err != nil {
		fmt.Println("Error getting secret from AWS Secret Manager: " + err.Error())
		return dataSecret, err
	}

	// Unmarshal the secret string into the dataSecret struct.
	// * is used to dereference the pointer to get the actual value of key.SecretString.
	json.Unmarshal([]byte(*key.SecretString), &dataSecret)
	fmt.Println("Secret from AWS Secret Manager: " + secretName)
	return dataSecret, nil
}
