package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	listenInterface      = ":8080"
	staticFilesDirectory = "./static"
)

func main() {
	staticFiles := http.FileServer(http.Dir(staticFilesDirectory))
	http.Handle("/", staticFiles)
	fmt.Println(fmt.Sprintf("Listening on http://localhost%v/index.html", listenInterface))
	err := http.ListenAndServe(listenInterface, nil)
	if err != nil {
		log.Fatal(err)
	}
}
