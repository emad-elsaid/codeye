package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/emad-elsaid/codeye"
)

func main() {
	http.HandleFunc("/", codeye.Home)
	http.HandleFunc("/file_structure", codeye.FileStructure)
	http.HandleFunc("/flame_graph", codeye.FlameGraph)

	var port = flag.String("port", ":0", "Server port number")
	flag.Parse()

	listener, err := net.Listen("tcp", *port)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("http://localhost:%d", listener.Addr().(*net.TCPAddr).Port)
	log.Printf("Serving on: %s\n", url)
	go codeye.OpenBrowser(url)
	http.Serve(listener, nil)
}
