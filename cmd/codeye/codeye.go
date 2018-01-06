package main

import (
	"log"
	"net/http"

	"github.com/emad-elsaid/codeye/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Home)

	log.Println("Serving on port :3000")
	http.ListenAndServe(":3000", nil)
}
