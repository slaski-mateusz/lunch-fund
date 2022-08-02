package main

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
