package getPages

import (
	"html"

	"github.com/PuerkitoBio/goquery"
)

//サンデーうぇぶりの更新取得
func GetSundaywebry() ([]TodayUpdate, error) {
	doc, err := goquery.NewDocument(SundaywebryUrl)
	if err != nil {
		return nil, err
	}

	//最新の漫画を取得
	updates := []TodayUpdate{}
	doc.Find("section[class=" + "box2" + "]").Each(func(_ int, s *goquery.Selection) {
		update := TodayUpdate{}
		s.Find("a").Each(func(_ int, s2 *goquery.Selection) {
			update.Url, _ = s2.Attr("href")
			s2.Find("img").Each(func(_ int, s3 *goquery.Selection) {
				update.ImgUrl, _ = s3.Attr("src")
				update.Title, _ = s3.Attr("alt")
				update.Title = html.EscapeString(update.Title)
				updates = append(updates, update)
			})

		})
	})
	return updates, nil

}
