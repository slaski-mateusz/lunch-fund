package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"

	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var DbStorePath *string
var databases map[string]*sql.DB

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

var dbModQueries = struct {
	add1stAdminQ        string
	addMemberQ          string
	updateMemberQ       string
	deleteMemberQ       string
	listMembersQ        string
	addOrderQ           string
	updateOrderQ        string
	deleteOrderQ        string
	listOrdersQ         string
	addOrderDetailsQ    string
	updateOrderDetailsQ string
	deleteOrderDetailsQ string
	listOrdersDetailsQ  string
}{
	add1stAdminQ:        ``,
	addMemberQ:          ``,
	updateMemberQ:       ``,
	deleteMemberQ:       ``,
	listMembersQ:        `SELECT * FROM members;`,
	addOrderQ:           ``,
	updateOrderQ:        ``,
	deleteOrderQ:        ``,
	listOrdersQ:         `SELECT * FROM orders;`,
	addOrderDetailsQ:    ``,
	updateOrderDetailsQ: ``,
	deleteOrderDetailsQ: ``,
	listOrdersDetailsQ:  `SELECT * FROM orders_details;`,
}

type Money int32
type Timestamp int64

type Member struct {
	id     int64
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Active bool   `json:"active"`
}

type Order struct {
	id           int64
	Name         string
	Timestamp    int64
	FounderId    int64
	deliveryCost Money
	tipCost      Money
}

type OrderMember struct {
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

func TeamFilename(teamName string) string {
	return fmt.Sprintf("team_%v.db", teamName)
}

func DbPathWithName(dbPath string, dbName string) string {
	return strings.Join(
		[]string{
			dbPath,
			dbName,
		},
		string(os.PathSeparator),
	)
}

func DbExist(dbFilePath string) (bool, error) {
	_, err := os.Stat(dbFilePath)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func connectDB(teamName string) error {
	dbFilePath := DbPathWithName(
		*DbStorePath,
		TeamFilename(teamName),
	)
	if dbe, _ := DbExist(dbFilePath); dbe {
		var err error
		databases[teamName], err = sql.Open("sqlite3", dbFilePath)
		defer databases[teamName].Close()
		if err != nil {
			return err
		}
	}
	return nil

}

// Teams

func InitTeamDatabase(teamName string) error {
	if dbc, ok := databases[teamName]; ok {
		for _, initQuery := range dbInitQueries {
			dbCursor, errPre := dbc.Prepare(initQuery)
			if errPre != nil {
				return errPre
			}
			_, errExe := dbCursor.Exec()
			if errExe != nil {
				return errExe
			}
		}
	}
	return nil
}

func ListTeams() []string {
	files, err := ioutil.ReadDir(*DbStorePath)
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

func RenameTeam() error {
	// TODO: Reanaming Team
	return errors.New("Error when renaming team")
}

func DeleteTeam() error {
	// TODO: Deleting Team
	return errors.New("Error when deleting team")
}

// Members

func ListMembers(teamName string) ([]Member, error) {
	dbFilePath := DbPathWithName(
		*DbStorePath,
		TeamFilename(teamName),
	)
	dbe, errExist := DbExist(dbFilePath)
	if errExist != nil {
		return nil, errExist
	}
	if dbe {
		db, err := sql.Open("sqlite3", dbFilePath)
		defer db.Close()
		if err != nil {
			return nil, err
		}
		query := dbModQueries.listMembersQ
		dbCursor, errPre := db.Prepare(query)
		if errPre != nil {
			return nil, errPre
		}
		data, errExe := dbCursor.Exec()
		if errExe != nil {
			return nil, errExe
		}
		fmt.Println(data.RowsAffected())
		mm := []Member{}
		return mm, nil
	}
	return nil, errors.New("Unknown problem when getting members from database")
}

func AddMember(team string, newMember Member) error {
	// TODO: adding new member to database
	return errors.New("Unknown problem when adding member to database")
}

func UpdateMember(team string, memberData Member) error {
	// TODO: update member data
	return errors.New("Unknown problem when updating member in database")
}

func DeleteMember(team string, memberData Member) error {
	// TODO: delete member from team
	return errors.New("Unknown problem when delete member from database")
}

// Orders

func ListOrders(team string) error {
	// TODO: Listing orders
	return errors.New("Unknown problem when listing orders")
}

func AddOrder(team string, newOrder Order) error {
	// TODO: adding new order to database
	return errors.New("Unknown problem when adding order to database")
}

func AddOrder(team string, newOrder Order) error {
	// TODO: adding new order to database
	return errors.New("Unknown problem when adding order to database")
}

func AddOrder(team string, newOrder Order) error {
	// TODO: adding new order to database
	return errors.New("Unknown problem when adding order to database")
}
