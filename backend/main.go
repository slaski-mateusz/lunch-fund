package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/slaski-mateusz/lunch-fund/backend/api"
	"github.com/slaski-mateusz/lunch-fund/backend/db"
)

func RunServer(netIntf string, netPort uint) {
	// TODO embedd static files within application?
	api.ActivateNodeRoute("", api.RoutingStructure)
	api.Router.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(
				http.Dir("./static"),
			),
		),
	)
	webIntf := fmt.Sprintf("%v:%v", netIntf, netPort)
	log.Fatal(http.ListenAndServe(webIntf, api.Router))
}

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
		port: 8088,
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
	RunServer(
		srvInt.host,
		uint(srvInt.port),
	)
}
