package controllers

import "net/http"

func SetupRoutes() {
	http.HandleFunc("/", handleRoot)
}
