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

	"github.com/slaski-mateusz/lunch-fund/model"
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
		ConnectedDatabases[teamName], err = sql.Open(dbEngine, dbFilePath)
		defer ConnectedDatabases[teamName].Close()
		if err != nil {
			return err
		}
	}
	return nil

}

// Teams

func InitTeamDatabase(teamName string) error {
	dbFilePath := DbPathWithName(
		*DbStorePath,
		TeamFilename(teamName),
	)
	if dbe, _ := DbExist(dbFilePath); !dbe {
		database, err := sql.Open(dbEngine, dbFilePath)
		defer database.Close()
		if err == nil {
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
			return err
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

// Now done in API
// TODO: Move this here
// func RenameTeam() error {
// 	// TODO: Reanaming Team
// 	return errors.New("Error when renaming team")
// }

// func DeleteTeam() error {
// 	// TODO: Deleting Team
// 	return errors.New("Error when deleting team")
// }

// Members

func ListMembers(teamName string) ([]model.Member, error) {
	dbFilePath := DbPathWithName(
		*DbStorePath,
		TeamFilename(teamName),
	)
	dbe, errExist := DbExist(dbFilePath)
	if errExist != nil {
		return nil, errExist
	}
	if dbe {
		db, errOp := sql.Open(dbEngine, dbFilePath)
		defer db.Close()
		if errOp != nil {
			return nil, errOp
		}
		query := dbCrudQueries.listMembersQ
		dbCursor, errPre := db.Prepare(query)
		if errPre != nil {
			return nil, errPre
		}
		data, errExe := dbCursor.Query()
		defer dbCursor.Close()
		if errExe != nil {
			return nil, errExe
		}
		members := []model.Member{}
		for data.Next() {
			var (
				recid       int64
				recname     string
				recemail    string
				recphone    string
				recisadmin  bool
				recisactive bool
				recavatar   []byte
			)
			errNx := data.Scan(
				&recid,
				&recname,
				&recemail,
				&recphone,
				&recisactive,
				&recisadmin,
				&recavatar,
			)
			if errNx != nil {
				return nil, errNx
			}
			var recmemeber model.Member
			recmemeber.Id = recid
			recmemeber.MemberName = recname
			recmemeber.Email = recemail
			recmemeber.Phone = recphone
			recmemeber.Active = recisactive
			members = append(members, recmemeber)
		}
		return members, nil
	}
	return nil, errors.New("Unknown problem when getting members from database")
}

func AddMember(newMember model.TeamMember) error {
	// TODO: adding new member to database
	dbFilePath := DbPathWithName(
		*DbStorePath,
		TeamFilename(newMember.TeamName),
	)
	dbe, errExist := DbExist(dbFilePath)
	if errExist != nil {
		return errExist
	}
	if dbe {
		db, errOp := sql.Open(dbEngine, dbFilePath)
		defer db.Close()
		if errOp != nil {
			return errOp
		}
		_, errExe := db.Exec(
			dbCrudQueries.addMemberQ,
			newMember.MemberName,
			newMember.Email,
			newMember.Phone,
			false,
			true,
		)
		if errExe != nil {
			return errExe
		}
		return nil
	}
	return errors.New("Unknown problem when adding member to database")
}

func UpdateMember(team string, memberData model.Member) error {
	// TODO: update member data
	return errors.New("Unknown problem when updating member in database")
}

func DeleteMember(team string, memberData model.Member) error {
	// TODO: delete member from team
	return errors.New("Unknown problem when delete member from database")
}

// Orders

func ListOrders(team string) error {
	// TODO: Listing orders
	return errors.New("Unknown problem when listing orders")
}

func AddOrder(team string, newOrder model.Order) error {
	// TODO: adding new order to database
	return errors.New("Unknown problem when adding order to database")
}

func UpdateOrder(team string, newOrder model.Order) error {
	// TODO: adding new order to database
	return errors.New("Unknown problem when updating order in database")
}

func DeleteOrder(team string, newOrder model.Order) error {
	// TODO: adding new order to database
	return errors.New("Unknown problem when removing order from database")
}
