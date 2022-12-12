package main

import (
	"log"
	"net/http"
)

var (
	listenInterface = ":8080"
	servedDir       = "."
)

func main() {
	log.Fatal(
		http.ListenAndServe(
			*&listenInterface,
			http.FileServer(
				http.Dir(*&servedDir),
			),
		),
	)
}
