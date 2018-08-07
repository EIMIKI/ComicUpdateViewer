package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DbUser struct {
	Name string
	Pass string
}

var Conn *sql.DB

func getDbUserEnv() DbUser {
	dbUser := DbUser{}
	dbUser.Name = os.Getenv("DBUSER")
	dbUser.Pass = os.Getenv("DBPASS")
	return dbUser

}
func Connect() (err error) {
	dbUser := getDbUserEnv()
	conn, err := sql.Open("mysql", dbUser.Name+":"+dbUser.Pass+"@/cuv")
	if err != nil {
		return err
	}

	Conn = conn
	return nil
}
