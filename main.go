package main

import (
	"ComicUpdateViewer/api"
	"ComicUpdateViewer/db"
	"ComicUpdateViewer/getPages"
	"time"

	"github.com/gin-gonic/gin"
)

func check() {
	for {
		updates, err := getPages.GetPages()
		if err != nil {
			panic(err)
		}

		for _, update := range updates {
			err = db.Push(update)
		}
		if err != nil {
			panic(err)
		}
		time.Sleep(12 * time.Hour)
	}
}

func main() {

	r := gin.Default()
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Conn.Close()

	r.GET("/api/getToday", api.GetToday)
	go check()
}
