/*
Package controllers is a middle layer that lets us set http routes
*/
package controllers

import (
	"net/http"
)

func SetupRoutes() {
	http.Handle("/ext/", http.StripPrefix("/ext/", http.FileServer(http.Dir("ext"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/about", HandleAbout)
	http.HandleFunc("/", HandleRoot)
}
