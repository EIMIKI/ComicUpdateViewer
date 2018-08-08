package view

import (
	"ComicUpdateViewer/db"

	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

type ComicData struct {
	Today []db.Result
	Past  []db.Result
}

// すべての更新データを含めたHTMLを提供する
func ComicList(c *gin.Context) {
	date := db.GetDate()
	data, err := db.GetPastUpdate("select comic_id,title,url,img from comics where date='" + date + "'")
	if err != nil {
		log.Fatalln(err)
	}

	comicData := ComicData{Today: data}
	data, err = db.GetPastUpdate("select comic_id,title,url,img from comics where not(date = '" + date + "')")
	if err != nil {
		log.Fatalln(err)
	}
	comicData.Past = data

	c.HTML(http.StatusOK, "comicList.html", comicData)

}
