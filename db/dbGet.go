package db

import (
	"fmt"
)

type Result struct {
	ComicId string
	Title   string
	Url     string
	Img     string
}

// DBからsqlQueryに沿ったデータを持ってくる
func GetPastUpdate(sqlQuery string) ([]Result, error) {
	resultRows, err := Conn.Query(sqlQuery)
	//resultRows, err := conn.Query("select comic_id,title,url,img from comics where date='" + date + "'")
	if err != nil {
		return nil, err
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
