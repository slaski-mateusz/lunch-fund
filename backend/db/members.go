package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func ListMembers(teamName string) ([]model.Member, error) {
	errCon := connectDB(teamName)
	if errCon != nil {
		return nil, errCon
	} else {
		dbinst := ConnectedDatabases[teamName]
		query := dbCrudQueries.membersListQ
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
			var recmemeber model.Member
			var secret sql.NullString
			errNx := data.Scan(
				&recmemeber.Id,
				&recmemeber.MemberName,
				&recmemeber.Email,
				&recmemeber.Phone,
				&recmemeber.IsAdmin,
				&recmemeber.IsActive,
				&secret,
			)
			if errNx != nil {
				return nil, errNx
			}
			members = append(members, recmemeber)
		}
		return members, nil
	}
}

func AddMember(newMember model.Member) error {
	errCon := connectDB(newMember.TeamName)
	if errCon != nil {
		return errCon
	} else {
		dbinst := ConnectedDatabases[newMember.TeamName]
		_, errExe := dbinst.Exec(
			dbCrudQueries.memberAddQ,
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
}

func UpdateMember(memberData model.Member) error {
	errCon := connectDB((memberData.TeamName))
	if errCon != nil {
		return errCon
	} else {
		dbinst := ConnectedDatabases[memberData.TeamName]
		_, errExe := dbinst.Exec(
			dbCrudQueries.memberUpdateQ,
			memberData.MemberName,
			memberData.Email,
			memberData.Phone,
			memberData.IsAdmin,
			memberData.IsActive,
			memberData.Id,
		)
		if errExe != nil {
			return errExe
		}
		return nil
	}
}

func DeleteMember(deleteMember model.Member) error {
	errCon := connectDB(deleteMember.TeamName)
	if errCon != nil {
		return errCon
	} else {
		dbinst := ConnectedDatabases[deleteMember.TeamName]
		row := dbinst.QueryRow(
			dbCrudQueries.memberCheckIfExistQ,
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
				dbCrudQueries.memberDeleteMQ,
				deleteMember.Id,
			)
			if errExe != nil {
				return errExe
			}
			return nil
		}
	}
}
