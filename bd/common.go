package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mauro-arcelles/gambituser/models"
	"github.com/mauro-arcelles/gambituser/secretm"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))

	return err
}

func DbConnect() error {
	Db, err := sql.Open("mysql", ConnStr(SecretModel))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conectado a la BD")

	return nil
}

func ConnStr(clave models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = clave.Username
	authToken = clave.Password
	dbEndpoint = clave.Host
	dbName = "gambit"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)

	return dsn
}
