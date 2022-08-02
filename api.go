package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

type newTeam struct {
	TeamName  string `json:"team_name"`
	AdminName string `json:"admin_name"`
}

type renameTeam struct {
	OldTeamName string `json:"old_team_name"`
	NewTeamName string `json:"new_team_name"`
}

type teamToDel struct {
	TeamName string `json:"team_name"`
}

func docpageHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("docpageHandler called")
}

type MemberApi struct {
	Team string `json:"team"`
	Member
}

func membersHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("membersHandler called")
	var newMembererData MemberApi
	switch requ.Method {
	case "PUT":
		errDecode := json.NewDecoder(requ.Body).Decode(&newMembererData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		} else {
			fmt.Println(fmt.Sprintf("Received new member data: %v", newMembererData))
		}
	}
}

func ordersHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("ordersHandler called")
}

func debtsHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("debtsHandler called")
}

func teamsHandler(resWri http.ResponseWriter, requ *http.Request) {
	var newTeamData newTeam
	fmt.Println("teamsHandler called")
	switch requ.Method {
	case "PUT":
		errDecode := json.NewDecoder(requ.Body).Decode(&newTeamData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		} else {
			errInit := initTeamDatabase(*dbStorePath, newTeamData.TeamName)
			if errInit != nil {
				http.Error(
					resWri,
					errInit.Error(),
					http.StatusBadRequest,
				)
				return
			}
		}
	case "POST":
		var renameTeamData renameTeam
		errDecode := json.NewDecoder(requ.Body).Decode(&renameTeamData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		} else {
			oldTeamNamePath := dbPathWithName(
				*dbStorePath,
				teamFilename(renameTeamData.OldTeamName),
			)
			newTeamNamePath := dbPathWithName(
				*dbStorePath,
				teamFilename(renameTeamData.NewTeamName),
			)
			oldTeamPathExist, errCheckExist := dbExist(oldTeamNamePath)
			if errCheckExist != nil {
				http.Error(
					resWri,
					errDecode.Error(),
					http.StatusBadRequest,
				)
				return
			} else {
				if oldTeamPathExist {
					errRename := os.Rename(
						oldTeamNamePath,
						newTeamNamePath,
					)
					if errRename != nil {
						http.Error(
							resWri,
							errDecode.Error(),
							http.StatusBadRequest,
						)
						return
					}
				} else {
					http.Error(
						resWri,
						fmt.Sprintf("There is no file '%v' to rename", oldTeamNamePath),
						http.StatusBadRequest,
					)
					return
				}
			}
		}
	case "DELETE":
		var teamToDelData teamToDel
		errDecode := json.NewDecoder(requ.Body).Decode(&teamToDelData)
		if errDecode != nil {
			http.Error(
				resWri,
				errDecode.Error(),
				http.StatusBadRequest,
			)
			return
		} else {
			teamToDelPath := dbPathWithName(
				*dbStorePath,
				teamFilename(teamToDelData.TeamName),
			)
			dbToDelExist, errCheckExist := dbExist(teamToDelPath)
			if errCheckExist != nil {
				http.Error(
					resWri,
					errDecode.Error(),
					http.StatusBadRequest,
				)
				return
			} else {
				if dbToDelExist {
					errRemove := os.Remove(teamToDelPath)
					if errRemove != nil {
						http.Error(
							resWri,
							errDecode.Error(),
							http.StatusBadRequest,
						)
						return
					}
				} else {
					http.Error(
						resWri,
						fmt.Sprintf("There is no file '%v' to delete", teamToDelPath),
						http.StatusBadRequest,
					)
					return
				}

			}
		}
	default:
		var teamsNames []string
		for _, teamFilename := range listTeams(*dbStorePath) {
			teamName := strings.TrimSuffix(
				strings.TrimPrefix(
					teamFilename,
					"team_"),
				".db",
			)
			teamsNames = append(teamsNames, teamName)
		}
		json.NewEncoder(resWri).Encode(
			teamsNames,
		)
		return
	}
}

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

func handleRequests(net_intf string, net_port uint) {
	activateApiNode("", apiStructure)
	webIntf := fmt.Sprintf("%v:%v", net_intf, net_port)
	log.Fatal(http.ListenAndServe(webIntf, router))
}
