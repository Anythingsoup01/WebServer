package controllers

import (
	"html/template"
	"net/http"

	"github.com/Anythingsoup01/WebServer/models"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/about.html"))
	data := models.GetAboutData()
	err := tpl.Execute(w, data);
	if err != nil {
		println("Failed to execute HandleAbout::template");
	}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/index.html"))
	data := models.GetRootData()
	err := tpl.Execute(w, data);
	if err != nil {
		println("Failed to execute HandleRoot::template");
	}
}
