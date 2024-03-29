package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbEngine = "sqlite3"
)

var DbStorePath *string
var ConnectedDatabases map[string]*sql.DB

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
	return false, err
}

func connectDB(teamName string) error {
	dbFilePath := DbPathWithName(
		*DbStorePath,
		TeamFilename(teamName),
	)
	if dbe, _ := DbExist(dbFilePath); dbe {
		if ConnectedDatabases == nil {
			ConnectedDatabases = make(map[string]*sql.DB)
		}
		if ConnectedDatabases[teamName] == nil {
			var err error
			ConnectedDatabases[teamName], err = sql.Open(
				dbEngine,
				"file:"+dbFilePath+"?_foreign_keys=true",
			)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New(
			fmt.Sprintf(
				"No database file '%v' for team '%v'",
				dbFilePath,
				teamName,
			),
		)
	}
	return nil

}
