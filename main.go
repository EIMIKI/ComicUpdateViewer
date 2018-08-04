package main

import (
	"ComicUpdateViewer/db"
	"ComicUpdateViewer/getPages"
	"fmt"
)

func main() {
	conn, err := db.Conn()
	if err != nil {
		panic(err)
	}
	fmt.Println(conn)
	getPages.GetPages()
}
