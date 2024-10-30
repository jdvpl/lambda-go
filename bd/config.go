package bd

import (
	"jdvpl/models"
	"jdvpl/secretManager"
	"os"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretManager.GetSecret(os.Getenv("SECRETNAME"))
	if err != nil {
		return err
	}
	return nil
}
