package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/slaski-mateusz/lunch-fund/backend/db"
	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func teamsHandler(resWri http.ResponseWriter, requ *http.Request) {
	switch requ.Method {
	case "PUT":
		var newTeamData model.Team
		errDecode := json.NewDecoder(requ.Body).Decode(&newTeamData)
		if errDecode != nil {
			http.Error(
				resWri,
				errDecode.Error(),
				http.StatusBadRequest,
			)
			return
		} else {
			if newTeamData.TeamName != "" {
				errInit := db.InitTeamDatabase(newTeamData.TeamName)
				if errInit != nil {
					http.Error(
						resWri,
						errInit.Error(),
						http.StatusBadRequest,
					)
					return
				}
			} else {
				http.Error(
					resWri,
					"Empty team name or no 'team' key in request body",
					http.StatusBadRequest,
				)
			}

		}
	case "POST":
		var renameTeamData model.TeamRename
		errDecode := json.NewDecoder(requ.Body).Decode(&renameTeamData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		} else {
			oldTeamNamePath := db.DbPathWithName(
				*db.DbStorePath,
				db.TeamFilename(renameTeamData.OldTeamName),
			)
			newTeamNamePath := db.DbPathWithName(
				*db.DbStorePath,
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
		var teamToDelData model.Team
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
				*db.DbStorePath,
				db.TeamFilename(teamToDelData.TeamName),
			)
			dbToDelExist, errCheckExist := db.DbExist(teamToDelPath)
			if errCheckExist != nil {
				http.Error(
					resWri,
					errCheckExist.Error(),
					http.StatusBadRequest,
				)
				return
			} else {
				if dbToDelExist {
					errRemove := os.Remove(teamToDelPath)
					if errRemove != nil {
						http.Error(
							resWri,
							errRemove.Error(),
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
	case "GET":
		var teamsNames []string
		for _, teamFilename := range db.ListTeams() {
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
