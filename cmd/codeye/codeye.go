package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/emad-elsaid/codeye/controllers"
	"github.com/emad-elsaid/codeye/utils"
)

func main() {
	http.HandleFunc("/", controllers.Home)

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	fmt.Println("Using port:")

	url := fmt.Sprintf("http://localhost:%d", listener.Addr().(*net.TCPAddr).Port)
	log.Printf("Serving on: %s\n", url)
	go utils.OpenBrowser(url)
	http.Serve(listener, nil)
}
