package getPages

import (
	"fmt"
)

type TodayUpdate struct {
	Url    string
	ImgUrl string
	Title  string
}

var (
	UrasundayUrl   = "http://urasunday.com/"
	SundaywebryUrl = "https://www.sunday-webry.com/"
)

func GetPages() error {
	UsUpdates, err := GetUrasunday()
	if err != nil {
		return err
	}
	SwUpdates, err := GetSundaywebry()
	if err != nil {
		return err
	}

	fmt.Println(UsUpdates)
	fmt.Println(SwUpdates)

	return nil
}
