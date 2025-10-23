package models

type ListData struct {
	Title string
	Items []string
}

func GetListData() ListData {
	return ListData {
		Title: "List Data",
		Items: []string {
			"First Item",
			"Second Item",
			"Third Item",
		},
	}
}
