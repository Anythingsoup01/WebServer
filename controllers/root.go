package controllers

import (
	"html/template"
	"net/http"
	"github.com/Anythingsoup01/WebServer/models"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("views/index.html"))
	data := models.GetListData()
	tpl.Execute(w, data);
}
