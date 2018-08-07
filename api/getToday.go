package api

import (
	"ComicUpdateViewer/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// アクセス時点の日付を持つデータをjsonで返す
func GetToday(c *gin.Context) {
	date := db.GetDate()
	data, err := db.GetPastUpdate("select comic_id,title,url,img from comics where date='" + date + "'")
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, data)
}

// アクセス時点の日付を持たないデータをjsonで返す
func GetPast(c *gin.Context) {
	date := db.GetDate()
	data, err := db.GetPastUpdate("select comic_id,title,url,img from comics where not(date = '" + date + "')")
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, data)
}
