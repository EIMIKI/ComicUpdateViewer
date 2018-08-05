package main

import (
	"ComicUpdateViewer/db"
	"ComicUpdateViewer/getPages"
	"database/sql"
)

func check(conn *sql.DB) {
	updates, err := getPages.GetPages()
	if err != nil {
		panic(err)
	}

	for _, update := range updates {
		err = db.Push(conn, update)
	}
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := db.Conn()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	check(conn)
}
