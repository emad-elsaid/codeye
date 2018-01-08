package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/emad-elsaid/codeye"
)

func main() {
	http.HandleFunc("/", codeye.Home)

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("http://localhost:%d", listener.Addr().(*net.TCPAddr).Port)
	log.Printf("Serving on: %s\n", url)
	go codeye.OpenBrowser(url)
	http.Serve(listener, nil)
}
