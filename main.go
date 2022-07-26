package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// SQLite3 database related declarations

var dbInitQueries = map[string]string{
	"activateForeginKeys": "PRAGMA foreign_keys = ON;",
	"createTableMembers": `CREATE TABLE IF NOT EXISTS members (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		phone TEXT NOT NULL UNIQUE,
		is_admin INTEGER NOT NULL DEFAULT 0,
		active INTEGER NOT NULL DEFAULT 1,
		avatar BLOB)`,
	"createTableOrders": `CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		timestamp INTEGER NOT NULL,
		founder_id INTEGER NOT NULL,
		delivery_cost INTEGER DEFAULT 0,
		tip_cost  INTEGER DEFAULT 0,
		FOREIGN KEY (founder_id) REFERENCES members (id) ON UPDATE CASCADE ON DELETE RESTRICT)`,
	"createTableOrdersDetails": `CREATE TABLE IF NOT EXISTS orders_details (
		order_id INTEGER NOT NULL,
		member_id INTEGER NOT NULL,
		amount INTEGER NOT NULL,
		FOREIGN KEY (order_id) REFERENCES orders (id) ON UPDATE CASCADE ON DELETE RESTRICT,
		FOREIGN KEY (member_id) REFERENCES members (id) ON UPDATE CASCADE ON DELETE RESTRICT)`,
	"createTableDebts": `CREATE TABLE IF NOT EXISTS debts (
		debtor_id INTEGER NOT NULL,
		creditor_id INTEGER NOT NULL,
		amount INTEGER NOT NULL,
		return_timestamp INTEGER,
		FOREIGN KEY (debtor_id) REFERENCES members (id) ON UPDATE CASCADE ON DELETE RESTRICT,
		FOREIGN KEY (creditor_id) REFERENCES members (id) ON UPDATE CASCADE ON DELETE RESTRICT)`,
}

var dbModQueries = map[string]string{
	"add1stAdmin":        ``,
	"addMember":          ``,
	"updateMember":       ``,
	"deleteMember":       ``,
	"listMembers":        `SELECT * FROM members;`,
	"addOrder":           ``,
	"updateOrder":        ``,
	"deleteOrder":        ``,
	"listOrders":         `SELECT * FROM orders;`,
	"addOrderDetails":    ``,
	"updateOrderDetails": ``,
	"deleteOrderDetails": ``,
	"listOrdersDetails":  `SELECT * FROM orders_details;`,
}

type Money int32
type Timestamp int64

type Member struct {
	id     int64
	Name   string
	Email  string
	Phone  string
	Active bool
}

type Order struct {
	id           int64
	Name         string
	Timestamp    int64
	FounderId    int64
	deliveryCost Money
	tipCost      Money
}

type OrderDetail struct {
	OrderId  int64
	MemberId int64
	Amount   Money
}

type Debt struct {
	DebtorId        int64
	CreditorId      int64
	Amount          Money
	ReturnTimestamp Timestamp
}

func teamFilename(teamName string) string {
	return fmt.Sprintf("team_%v.db", teamName)
}

func dbPathWithName(dbPath string, dbName string) string {
	return strings.Join(
		[]string{
			dbPath,
			dbName,
		},
		string(os.PathSeparator),
	)
}

func dbExist(dbFilePath string) (bool, error) {
	_, err := os.Stat(dbFilePath)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func initTeamDatabase(dbStorePath string, teamName string) error {
	// fmt.Println(dbFilePath)
	dbFilePath := dbPathWithName(
		dbStorePath,
		teamFilename(teamName),
	)
	if dbe, _ := dbExist(dbFilePath); !dbe {
		db, err := sql.Open("sqlite3", dbFilePath)
		defer db.Close()
		if err != nil {
			return err
		}
		for _, initQuery := range dbInitQueries {
			dbCursor, errPre := db.Prepare(initQuery)
			if errPre != nil {
				return errPre
			}
			_, errExe := dbCursor.Exec()
			if errExe != nil {
				return errExe
			}
		}

	} else {
		return errors.New(
			fmt.Sprintf("Database file '%v' already exist", dbFilePath),
		)
	}
	return nil
}

func listTeams(dbStorePath string) []string {
	files, err := ioutil.ReadDir(dbStorePath)
	if err != nil {
		log.Fatal(err)
	}
	var teamDatabases []string
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "team_") && strings.HasSuffix(f.Name(), ".db") {
			teamDatabases = append(teamDatabases, f.Name())
		}
	}
	return teamDatabases
}

// API related declarations

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

func membersHandler(resWri http.ResponseWriter, requ *http.Request) {
	fmt.Println("membersHandler called")
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

func activate_api_node(in_uri string, node *ApiNode) {
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
		activate_api_node(apiUri, child)
	}
}

func handle_requests(net_intf string, net_port uint) {
	activate_api_node("", apiStructure)
	webIntf := fmt.Sprintf("%v:%v", net_intf, net_port)
	log.Fatal(http.ListenAndServe(webIntf, router))
}

// --------------------------------

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
	fmt.Println(fmt.Sprintf("Using %v directory as storage", dbStoreInd))
	handle_requests("127.0.0.1", 8080)
}
