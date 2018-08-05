package getPages

type TodayUpdate struct {
	Url    string
	ImgUrl string
	Title  string
}

var (
	UrasundayUrl   = "http://urasunday.com/"
	SundaywebryUrl = "https://www.sunday-webry.com/"
)

func GetPages() ([][]TodayUpdate, error) {
	UsUpdates, err := GetUrasunday()
	if err != nil {
		return nil, err
	}
	SwUpdates, err := GetSundaywebry()
	if err != nil {
		return nil, err
	}
	updates := [][]TodayUpdate{UsUpdates, SwUpdates}

	return updates, nil
}
