package db

import (
	"ComicUpdateViewer/getPages"
	"database/sql"
	"fmt"
	"time"
)

func Update(conn *sql.DB, comic getPages.TodayUpdate, date string) error {
	_, err := conn.Exec("update comics set date=" + date + ",img='" + comic.ImgUrl + "' where title='" + comic.Title + "'")
	if err != nil {
		return err
	}
	return nil
}

func Add(conn *sql.DB, comic getPages.TodayUpdate, date string) error {
	_, err := conn.Exec("insert into comics (title,url,img,date) value ('" + comic.Title + "','" + comic.Url + "','" + comic.ImgUrl + "','" + date + "')")
	if err != nil {
		return err
	}
	return nil
}

func Push(conn *sql.DB, comics []getPages.TodayUpdate) error {
	var exists string

	now := time.Now()
	date := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())

	for _, comic := range comics {
		err := conn.QueryRow("select exists (select * from comics where title='" + comic.Title + "')").Scan(&exists)

		if err != nil {
			return err
		}

		if exists == "0" {
			Add(conn, comic, date)
		} else {
			Update(conn, comic, date)
		}
	}
	return nil
}
