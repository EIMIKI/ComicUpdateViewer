package main

import (
	"ComicUpdateViewer/db"
	"ComicUpdateViewer/getPages"
	"database/sql"
	"time"
)

func check(conn *sql.DB) {
	for {
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
		time.Sleep(12 * time.Hour)
	}
}

func main() {
	conn, err := db.Conn()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	go check(conn)
}
