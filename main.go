package main

import (
	"flag"
	"fmt"

	"github.com/slaski-mateusz/lunch-fund/api"
	"github.com/slaski-mateusz/lunch-fund/db"
)

func main() {
	db.DbStorePath = flag.String(
		"dbStorage",
		"",
		"Location of teams databases",
	)
	flag.Parse()
	dbStoreInd := "application local"
	if *db.DbStorePath == "" || *db.DbStorePath == "." {
		if *db.DbStorePath == "" {
			*db.DbStorePath = "."
		}
	} else {
		dbStoreInd = *db.DbStorePath
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
	api.HandleRequests(
		srvInt.host,
		uint(srvInt.port),
	)
}
