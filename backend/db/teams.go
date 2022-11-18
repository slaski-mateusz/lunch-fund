package db

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func InitTeamDatabase(teamName string) error {
	dbFilePath := DbPathWithName(
		*DbStorePath,
		TeamFilename(teamName),
	)
	if dbe, _ := DbExist(dbFilePath); !dbe {
		database, errOp := sql.Open(
			dbEngine,
			"file:"+dbFilePath+"?_foreign_keys=true",
		)
		defer database.Close()
		if errOp == nil {
			dbCursorP, errPreP := database.Prepare(activateForeginKeys)
			if errPreP != nil {
				return errPreP
			}
			_, errExeP := dbCursorP.Exec()
			if errExeP != nil {
				return errExeP
			}
			for _, initQuery := range dbInitQueries {
				dbCursor, errPre := database.Prepare(initQuery)
				if errPre != nil {
					return errPre
				}
				_, errExe := dbCursor.Exec()
				if errExe != nil {
					return errExe
				}
			}
		} else {
			return errOp
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

// TODO: Move this here
// Now done in API
// func RenameTeam() error {
// 	return errors.New("Error when renaming team")
// }
// func DeleteTeam() error {
// 	return errors.New("Error when deleting team")
// }

// Members

// Orders
