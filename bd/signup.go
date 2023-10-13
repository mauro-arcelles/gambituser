package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mauro-arcelles/gambituser/models"
	"github.com/mauro-arcelles/gambituser/tools"
)

func SignUp(sign models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	query := fmt.Sprintf(`
    INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%v' , '%v', '%v')`,
		sign.UserEmail,
		sign.UserUUID,
		tools.FechaMySQL(),
	)

	fmt.Println(query)

	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecucion exitosa")
	return nil
}
