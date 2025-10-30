package models

import "time"

type ListData struct {
	Title string
	Items []string
	Time time.Time
}

func GetRootData() ListData {
	return ListData {
		Title: "Root File",
		Items: []string {
			"First Item",
			"Second Item",
			"Third Item",
		},
		Time: time.Now(),
	}
}

func GetAboutData() ListData {
	return ListData {
		Title: "About Me",
		Items: []string {
			"Computer Science Major",
			"Information Systems Minor",
			"Expected graduation: 2029",
		},
		Time: time.Now(),
	}
}
