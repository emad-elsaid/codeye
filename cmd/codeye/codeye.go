package main

import (
	"log"
	"net/http"

	"github.com/emad-elsaid/codeye/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)

	log.Println("Serving on port :3000")
	http.ListenAndServe(":3000", nil)
}
