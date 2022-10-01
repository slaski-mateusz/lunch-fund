package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiNode struct {
	nodeName string
	function func(resWri http.ResponseWriter, requ *http.Request)
	children []*ApiNode
}

var apiStructure = &ApiNode{
	nodeName: "",
	function: docpageHandler,
	children: []*ApiNode{
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
}

// Doc page

func docpageHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("docpageHandler called")
}

// Debts

func debtsHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("debtsHandler called")
}

// API Router, Nodes nctivation and General Handler

var router = mux.NewRouter().StrictSlash(true)

func activateApiNode(in_uri string, node *ApiNode) {
	var nodeUri string
	if node.nodeName == "" {
		nodeUri = "/"
	} else {
		nodeUri = fmt.Sprintf("%v/", node.nodeName)
	}
	apiUri := in_uri + nodeUri
	if node.function != nil {
		router.HandleFunc(apiUri, node.function)
	}
	for _, child := range node.children {
		activateApiNode(apiUri, child)
	}
}

func HandleRequests(netIntf string, netPort uint) {
	activateApiNode("", apiStructure)
	webIntf := fmt.Sprintf("%v:%v", netIntf, netPort)
	log.Fatal(http.ListenAndServe(webIntf, router))
}
