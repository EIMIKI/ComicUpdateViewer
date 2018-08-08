package main

import (
	"ComicUpdateViewer/api"
	"ComicUpdateViewer/db"
	"ComicUpdateViewer/getPages"
	"time"

	"ComicUpdateViewer/view"

	"github.com/gin-gonic/gin"
)

//定期実行して新着comicを探す
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

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")

	r.GET("/api/getToday", api.GetToday)
	r.GET("api/getPast", api.GetPast)
	r.GET("/", view.ComicList)

	go check()

	r.Run()
}
