package bd

import (
	"database/sql"
	"fmt"
	"jdvpl/models"
	"jdvpl/secretManager"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var DB *sql.DB

func ReadSecret() error {
	SecretModel, err = secretManager.GetSecret(os.Getenv("SECRETNAME"))
	if err != nil {
		fmt.Println("Error reading secret" + err.Error())
		return err
	}
	return nil
}

func DbConect() error {

	DB, err = sql.Open("mysql", ConnSrting(SecretModel))
	if err != nil {
		fmt.Println("Error reading secret" + err.Error())
		return err
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Error ping" + err.Error())
		return err
	}
	fmt.Println("Connection success")
	return nil
}

func ConnSrting(key models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = key.Username
	authToken = key.Password
	dbEndpoint = key.Host
	dbName = os.Getenv("DB")
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dns)
	return dns
}
