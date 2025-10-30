package controllers

import "net/http"

func SetupRoutes() {
	http.HandleFunc("/about", HandleAbout)
	http.HandleFunc("/", HandleRoot)
}
