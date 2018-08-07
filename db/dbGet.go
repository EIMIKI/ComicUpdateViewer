package db

import (
	"database/sql"
	"fmt"
)

type Result struct {
	ComicId string
	Title   string
	Url     string
	Img     string
}

// DBから実行した時点での日付のついたデータを持ってくる
func GetTodayUpdate(conn *sql.DB) ([]Result, error) {
	date := GetDate()
	resultRows, err := conn.Query("select comic_id,title,url,img from comics where date='" + date + "'")
	if err != nil {
		return err
	}

	defer resultRows.Close()

	results := []Result{}
	for resultRows.Next() {
		result := Result{}

		err := resultRows.Scan(&(result.ComicId), &(result.Title), &result.Url, &result.Img)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	fmt.Println(results)

	return results, err
}
