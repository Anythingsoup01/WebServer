/*
Package models provides a set of functions to retrieve Data
pertaining to a file.
*/
package models

import (
	"time"
)


type DivType int

const (
	StaticText DivType = iota
	Button
)


type DivData struct {
	Description string
	// DivType is used to determine what kind of div we want.
	// If DivType == 0 (StaticText) we don't expect a link to be with it.
	DivType DivType
	ImageLink string
	SiteLink string
}

type NavData struct {
	Description string
	Link string
}

type ListData struct {
	Title string
	DivItems []DivData
	NavLinks []NavData
	FooterItems []DivData
	Time time.Time
}

func GetRootData() ListData {
	return ListData {
		Title: "Home",
		DivItems: []DivData {
			{ Description: "Hello, welcome to my website! Below is a list of my most notable projects!", DivType: 0 },
			{ Description: "CatEngine", DivType: 1, ImageLink: "static/image/game_engine.png", SiteLink: "https://www.github.com/Anythingsoup01/CatEngine-CPP" },
			{ Description: "Possum-IDE", DivType: 1, ImageLink: "static/image/possum_ide.png", SiteLink: "https://www.github.com/Anythingsoup01/Possum-IDE" },
			{ Description: "Capybara", DivType: 1, ImageLink: "static/image/capybara.png", SiteLink: "https://www.github.com/Anythingsoup01/Capybara" },
			{ Description: "Ferret", DivType: 1, ImageLink: "static/image/ferret.png", SiteLink: "https://www.github.com/Anythingsoup01/Ferret" },
		},
		NavLinks: []NavData {
			{ Description: "About", Link: "about"},
		},
		FooterItems: []DivData {
			{ Description: "GitHub", DivType: 1, ImageLink: "static/image/github_icon.png", SiteLink: "https://www.github.com/Anythingsoup01" },
			{ Description: "LinkedIn", DivType: 1, ImageLink: "static/image/linkedin_icon.png", SiteLink: "https://www.linkedin.com/in/william-english-377242362/" },
		},
		Time: time.Now(),
	}
}

func GetAboutData() ListData {
	return ListData {
		Title: "About Me",
		DivItems: []DivData {
			{ Description: "Computer Science & Information Systems Major; Expected graduation: 2029", DivType: 0 },
			{ Description: "I've been programming since elementary, what originally started as a want to make games, quickly turned into wanting to make my own tools.", DivType: 0 },
			{ Description: "My first ambitious project I ever took on was making a game engine, to get a better understanding of C/C++, after a few years I swapped to linux, and daily drive it to this day, causing me to swap it over to linux.", DivType: 0 },
			{ Description: "My second ambitious project, however, isn't Ferret or Possum, it's Capybara, a tool I use in my game engine for hot reloading code, as a part of my ScriptEngine, I took heavy inspiration from Mono, but made it my own.", DivType: 0 },
		},
		NavLinks: []NavData {
			{ Description: "Home", Link: "/" },
		},
		Time: time.Now(),
	}
}
