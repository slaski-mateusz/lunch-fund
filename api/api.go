package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"

	"github.com/slaski-mateusz/lunch-fund/db"
	"github.com/slaski-mateusz/lunch-fund/model"
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

// Teams types

func docpageHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("docpageHandler called")
}

// Members

func membersHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("membersHandler called")

	switch requ.Method {
	case "GET":
		var teamData model.Team
		errDecode := json.NewDecoder(requ.Body).Decode(&teamData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		} else {
			fmt.Println(fmt.Sprintf("Received request for members of: %v", teamData))
			members, errMb := db.ListMembers(teamData.Name)
			if errMb != nil {
				http.Error(resWri, errMb.Error(), http.StatusBadRequest)
				return
			}
			fmt.Println(members)
			resWri.Write([]byte(fmt.Sprintf("%v", members)))
		}
	case "PUT":
		var newMembererData model.Member
		errDecode := json.NewDecoder(requ.Body).Decode(&newMembererData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		} else {
			fmt.Println(fmt.Sprintf("Received new member data: %v", newMembererData))
			// TODO: adding member to database
		}
	}
}

// Orders

func ordersHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("ordersHandler called")
}

// Debts

func debtsHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("debtsHandler called")
}

// Teams

func teamsHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("teamsHandler called")
	switch requ.Method {
	case "PUT":
		var newTeamData model.Team
		errDecode := json.NewDecoder(requ.Body).Decode(&newTeamData)
		if errDecode != nil {
			http.Error(
				resWri,
				errDecode.Error(),
				http.StatusBadRequest,
			)
			return
		} else {
			if newTeamData.Name != "" {
				errInit := db.InitTeamDatabase(newTeamData.Name)
				if errInit != nil {
					http.Error(
						resWri,
						errInit.Error(),
						http.StatusBadRequest,
					)
					return
				}
			} else {
				http.Error(
					resWri,
					"Empty team name or no 'team' key in request body",
					http.StatusBadRequest,
				)
			}

		}
	case "POST":
		var renameTeamData model.TeamRename
		errDecode := json.NewDecoder(requ.Body).Decode(&renameTeamData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		} else {
			oldTeamNamePath := db.DbPathWithName(
				*db.DbStorePath,
				db.TeamFilename(renameTeamData.OldTeamName),
			)
			newTeamNamePath := db.DbPathWithName(
				*db.DbStorePath,
				db.TeamFilename(renameTeamData.NewTeamName),
			)
			oldTeamPathExist, errCheckExist := db.DbExist(oldTeamNamePath)
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
		var teamToDelData model.Team
		errDecode := json.NewDecoder(requ.Body).Decode(&teamToDelData)
		if errDecode != nil {
			http.Error(
				resWri,
				errDecode.Error(),
				http.StatusBadRequest,
			)
			return
		} else {
			teamToDelPath := db.DbPathWithName(
				*db.DbStorePath,
				db.TeamFilename(teamToDelData.Name),
			)
			dbToDelExist, errCheckExist := db.DbExist(teamToDelPath)
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
	case "GET":
		var teamsNames []string
		for _, teamFilename := range db.ListTeams() {
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
