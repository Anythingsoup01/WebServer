package controllers

import (
	"html/template"
	"net/http"
	"github.com/Anythingsoup01/WebServer/models"
)

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/about.html"))
	data := models.GetAboutData()
	tpl.Execute(w, data);
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/index.html"))
	data := models.GetRootData()
	tpl.Execute(w, data);
}
