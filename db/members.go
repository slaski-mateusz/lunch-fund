package db

import (
	"database/sql"
	"errors"

	// "fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/slaski-mateusz/lunch-fund/model"
)

func ListMembers(teamName string) ([]model.Member, error) {
	errCon := connectDB(teamName)
	if errCon == nil {
		dbinst := ConnectedDatabases[teamName]
		query := dbCrudQueries.listMembersQ
		dbCursor, errPre := dbinst.Prepare(query)
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
	errCon := connectDB(newMember.TeamName)
	if errCon == nil {
		dbinst := ConnectedDatabases[newMember.TeamName]
		_, errExe := dbinst.Exec(
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

func DeleteMember(deleteMember model.TeamMember) error {
	errCon := connectDB(deleteMember.TeamName)
	if errCon == nil {
		dbinst := ConnectedDatabases[deleteMember.TeamName]
		row := dbinst.QueryRow(
			dbCrudQueries.checkIfMemberExistQ,
			deleteMember.Id,
		)
		errQuer := row.Scan(&deleteMember.Id)
		if errQuer != nil {
			if errQuer == sql.ErrNoRows {
				return errors.New("No such user in database")
			} else {
				return errQuer
			}
		} else {
			_, errExe := dbinst.Exec(
				dbCrudQueries.deleteMemberQ,
				deleteMember.Id,
			)
			if errExe != nil {
				return errExe
			}
			return nil
		}
	}
	return errors.New("Unknown problem when delete member from database")
}
