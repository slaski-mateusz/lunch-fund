package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/slaski-mateusz/lunch-fund/model"
)

func ListMembers(teamName string) ([]model.Member, error) {
	dbFilePath := DbPathWithName(
		*DbStorePath,
		TeamFilename(teamName),
	)
	fmt.Println(dbFilePath)
	dbe, errExist := DbExist(dbFilePath)
	if errExist != nil {
		fmt.Println("Nie ma bazy")
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

func DeleteMember(deleteMember model.TeamMember) error {
	dbFilePath := DbPathWithName(
		*DbStorePath,
		TeamFilename(deleteMember.TeamName),
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
		row := db.QueryRow(
			dbCrudQueries.checkIfMemberExistQ,
			deleteMember.Id,
		)
		errQuer := row.Scan(&deleteMember.Id)
		fmt.Println(errQuer)
		if errQuer != nil {
			if errQuer == sql.ErrNoRows {
				return errors.New("No such user in database")
			} else {
				return errQuer
			}
		} else {
			_, errExe := db.Exec(
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
