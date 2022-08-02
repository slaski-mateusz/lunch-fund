package main

import (
	"flag"
	"fmt"
)

var dbStorePath *string

func main() {
	dbStorePath = flag.String(
		"dbStorage",
		"",
		"Location of teams databases",
	)
	flag.Parse()
	dbStoreInd := "application local"
	if *dbStorePath == "" || *dbStorePath == "." {
		if *dbStorePath == "" {
			*dbStorePath = "."
		}
	} else {
		dbStoreInd = *dbStorePath
	}
	srvInt := struct {
		host string
		port int
	}{
		host: "127.0.0.1",
		port: 8080,
	}
	fmt.Println(fmt.Sprintf(
		"Using %v directory as storage",
		dbStoreInd,
	))
	fmt.Println(fmt.Sprintf(
		"Starting http server at %v:%v",
		srvInt.host,
		srvInt.port,
	))
	handleRequests(
		srvInt.host,
		uint(srvInt.port),
	)
}
