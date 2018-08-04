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

func getDbUserEnv() DbUser {
	dbUser := DbUser{}
	dbUser.Name = os.Getenv("DBUSER")
	dbUser.Pass = os.Getenv("DBPASS")
	return dbUser

}
func Conn() (*sql.DB, error) {
	dbUser := getDbUserEnv()
	conn, err := sql.Open("mysql", dbUser.Name+":"+dbUser.Pass+"@/cuv")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
