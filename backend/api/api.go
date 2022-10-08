package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteNode struct {
	nodeName string
	function func(resWri http.ResponseWriter, requ *http.Request)
	children []*RouteNode
}

var RoutingStructure = &RouteNode{
	nodeName: "",
	function: docpageHandler,
	children: []*RouteNode{
		// {
		// 	nodeName: "static",
		// 	function: staticFilesServe,
		// },
		{
			nodeName: "api",
			function: docpageHandler,
			children: []*RouteNode{
				{
					nodeName: "members",
					function: membersHandler,
				},
				{
					nodeName: "orders",
					function: ordersHandler,
				},
				{
					nodeName: "debts",
					function: debtsHandler,
				},
				{
					nodeName: "teams",
					function: teamsHandler,
				},
			},
		},
	},
}

// Doc page

func docpageHandler(resWri http.ResponseWriter, requ *http.Request) {
	resWri.Write([]byte("Not implemented yet"))
}

// Debts
// TODO Move to separate file
func debtsHandler(resWri http.ResponseWriter, requ *http.Request) {
	resWri.Write([]byte("Not implemented yet"))
}

// API Router, Nodes activation and General Handler

var Router = mux.NewRouter().StrictSlash(true)

func ActivateNodeRoute(in_uri string, node *RouteNode) {
	var nodeUri string
	if node.nodeName == "" {
		nodeUri = "/"
	} else {
		nodeUri = fmt.Sprintf("%v/", node.nodeName)
	}
	apiUri := in_uri + nodeUri
	if node.function != nil {
		Router.HandleFunc(apiUri, node.function)
	}
	for _, child := range node.children {
		ActivateNodeRoute(apiUri, child)
	}
}
