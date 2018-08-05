package getPages

import (
	"html"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// 元サイトが相対Pathでリンクされているため、絶対Pathに変換
func toAbsPath(url string) string {
	url = strings.Replace(url, "./", "", -1)
	return UrasundayUrl + url
}

// style属性を拾っていると余計な文字列が入るため、削除
func parseUrl(url string) string {
	url = strings.Replace(url, "background: url(", "", -1)
	url = strings.Replace(url, ")", "", -1)
	return toAbsPath(url)
}

//裏サンデーの更新取得
func GetUrasunday() ([]TodayUpdate, error) {
	doc, err := goquery.NewDocument(UrasundayUrl)
	if err != nil {
		return nil, err
	}

	//最新の漫画を取得
	updates := []TodayUpdate{}
	doc.Find("div[class=" + "indexMainImagesNew" + "]").Each(func(_ int, s *goquery.Selection) {
		update := TodayUpdate{}
		update.ImgUrl, _ = s.Attr("style")
		update.ImgUrl = parseUrl(update.ImgUrl)
		s.Find("a").Each(func(_ int, s2 *goquery.Selection) {
			update.Title = s2.Text()
			update.Title = html.EscapeString(update.Title)
			update.Url, _ = s2.Attr("href")
			update.Url = toAbsPath(update.Url)
			updates = append(updates, update)
		})
	})
	return updates, nil
}
