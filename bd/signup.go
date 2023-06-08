package bd

import (
	"fmt"
	"gambituser/models"
	"gambituser/tools"

	_ "github.com/go-sql-driver/mysql" // Driver de MySQL para el paquete database/sql
)

func SignUp(sig models.SignUp) error {
	fmt.Println("start registration")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	queryinsert := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "', '" + sig.UserUUID + "', '" + tools.DateMySQL() + "')"
	fmt.Println(queryinsert)

	_, err = Db.Exec(queryinsert)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	fmt.Println("SignUp > successful")
	return nil
}
