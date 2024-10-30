package bd

import (
	"fmt"
	"jdvpl/models"
	"jdvpl/tools"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.Signup) error {
	fmt.Println("start registration")
	err := DbConect()
	if err != nil {
		return err
	}
	defer DB.Close()
	sentence := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "', '" + sig.UserEmail + "', '" + tools.DateMysql() + "')"
	fmt.Println(sentence)
	_, err = DB.Exec(sentence)

	if err != nil {
		fmt.Println("Error insert" + err.Error())
		return err
	}
	return nil
}
