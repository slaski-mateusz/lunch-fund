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

type teamRequest struct {
	Team string `json:"team"`
}

type updateMember struct {
	teamRequest
	db.Member
}

type addMember updateMember

func membersHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("membersHandler called")

	switch requ.Method {
	case "PUT":
		var newMembererData updateMember
		errDecode := json.NewDecoder(requ.Body).Decode(&newMembererData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		} else {
			fmt.Println(fmt.Sprintf("Received new member data: %v", newMembererData))
			// TODO: adding member
		}
	case "GET":
		var tr teamRequest
		errDecode := json.NewDecoder(requ.Body).Decode(&tr)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		} else {
			fmt.Println(fmt.Sprintf("Received request for members of: %v", tr))
			mb, errMb := db.ListMembers(tr.Team)
			if errMb != nil {
				http.Error(resWri, errMb.Error(), http.StatusBadRequest)
				return
			}
			resWri.Write([]byte(fmt.Sprintf("%v", mb)))
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
			errInit := db.InitTeamDatabase(newTeamData.TeamName)
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
			oldTeamNamePath := db.DbPathWithName(
				*dbStorePath,
				db.TeamFilename(renameTeamData.OldTeamName),
			)
			newTeamNamePath := db.DbPathWithName()(
				*dbStorePath,
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
			teamToDelPath := db.DbPathWithName(
				*dbStorePath,
				db.TeamFilename(teamToDelData.TeamName),
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
	default:
		var teamsNames []string
		for _, teamFilename := range db.ListTeams(*dbStorePath) {
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

func HandleRequests(netIntf string, netPort uint) {
	activateApiNode("", apiStructure)
	webIntf := fmt.Sprintf("%v:%v", netIntf, netPort)
	log.Fatal(http.ListenAndServe(webIntf, router))
}
