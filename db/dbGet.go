package db

import (
	"fmt"
)

type Result struct {
	ComicId string `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Img     string `json:"img"`
}

// DBからsqlQueryに沿ったデータを持ってくる
func GetPastUpdate(sqlQuery string) ([]Result, error) {
	resultRows, err := Conn.Query(sqlQuery)
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
