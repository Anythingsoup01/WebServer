package main

import (
	"flag"
	"fmt"
	"net/http"
	"github.com/Anythingsoup01/WebServer/controllers"
)

func main() {
	flag.Parse()
	controllers.SetupRoutes()
	fmt.Println("Listening on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
