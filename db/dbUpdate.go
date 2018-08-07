package db

import (
	"ComicUpdateViewer/getPages"
	"fmt"
	"time"
)

// 実行時点での日付を返す
func GetDate() string {
	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	return date

}

// 既存のcomicデータを更新する
func Update(comic getPages.TodayUpdate, date string) error {
	_, err := Conn.Exec("update comics set date=" + date + ",img='" + comic.ImgUrl + "' where title='" + comic.Title + "'")
	if err != nil {
		return err
	}
	return nil
}

// 新規にcomicデータを追加する
func Add(comic getPages.TodayUpdate, date string) error {
	_, err := Conn.Exec("insert into comics (title,url,img,date) value ('" + comic.Title + "','" + comic.Url + "','" + comic.ImgUrl + "','" + date + "')")
	if err != nil {
		return err
	}
	return nil
}

// 当日分のcomicデータをデータベースへ送る
func Push(comics []getPages.TodayUpdate) error {
	var exists string

	date := GetDate()

	for _, comic := range comics {
		err := Conn.QueryRow("select exists (select * from comics where title='" + comic.Title + "')").Scan(&exists)

		if err != nil {
			return err
		}

		if exists == "0" {
			Add(comic, date)
		} else {
			Update(comic, date)
		}
	}
	return nil
}
