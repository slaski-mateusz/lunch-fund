package api

import (
	"encoding/json"
	"net/http"

	"github.com/slaski-mateusz/lunch-fund/backend/db"
	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func membersHandler(resWri http.ResponseWriter, requ *http.Request) {
	switch requ.Method {
	case http.MethodGet:
		var teamData model.Team
		errDecode := json.NewDecoder(requ.Body).Decode(&teamData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		}
		if teamData.TeamName == "" {
			http.Error(resWri, "Team name can not to be emty!\n", http.StatusBadRequest)
			return
		}
		members, errMb := db.ListMembers(teamData.TeamName)
		if errMb != nil {
			http.Error(resWri, errMb.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(resWri).Encode(members)
		return

	case http.MethodPut:
		var newMembererData model.Member
		errDecode := json.NewDecoder(requ.Body).Decode(&newMembererData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		}
		if newMembererData.TeamName == "" {
			http.Error(resWri, "Team name can not to be emty!\n", http.StatusBadRequest)
			return
		}
		errAdd := db.AddMember(newMembererData)
		if errAdd != nil {
			http.Error(resWri, errAdd.Error(), http.StatusBadRequest)
			return
		}
		resWri.Write([]byte("Member added\n"))
		return

	case http.MethodPost:
		var updatedMembererData model.Member
		errDecode := json.NewDecoder(requ.Body).Decode(&updatedMembererData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		}
		if updatedMembererData.TeamName == "" {
			http.Error(resWri, "Team name can not to be emty!\n", http.StatusBadRequest)
			return
		}
		errUpd := db.UpdateMember(updatedMembererData)
		if errUpd != nil {
			http.Error(resWri, errUpd.Error(), http.StatusBadRequest)
			return
		}
		resWri.Write([]byte("Member updated\n"))
		return

	case http.MethodDelete:
		var deletedMenberData model.Member
		errDecode := json.NewDecoder(requ.Body).Decode(&deletedMenberData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		}
		if deletedMenberData.TeamName == "" {
			http.Error(resWri, "Team name can not to be emty!\n", http.StatusBadRequest)
			return
		}
		errDel := db.DeleteMember(deletedMenberData)
		if errDel != nil {
			http.Error(resWri, errDel.Error(), http.StatusBadRequest)
			return
		}
		resWri.Write([]byte("Member deleted\n"))
		return
	}
}
