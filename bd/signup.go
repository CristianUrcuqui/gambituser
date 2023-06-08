package bd

import (
	"fmt"
	"gambituser/models"
	"gambituser/tools"

	_ "github.com/go-sql-driver/mysql" // Driver de MySQL para el paquete database/sql
)

// SignUp es una función que inserta un nuevo usuario en la base de datos.
// Acepta como parámetro una variable 'sig' de tipo SignUp (definido en el paquete models).
// Retorna un error si ocurre alguno durante el proceso.
func SignUp(sig models.SignUp) error {
	fmt.Println("start registration")

	// Intenta conectar a la base de datos
	err := DbConnect()
	if err != nil {
		return err
	}

	// Asegura que la conexión a la base de datos se cierre al finalizar la función
	defer Db.Close()

	// Prepara la consulta SQL
	stmt, err := Db.Prepare("INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Ejecuta la consulta SQL con los valores proporcionados
	_, err = stmt.Exec(sig.UserEmail, sig.UserUUID, tools.DateMySQL())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > successful")
	return nil
}
