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
		}
		members, errMb := db.ListMembers(teamData.TeamName)
		if errMb != nil {
			http.Error(resWri, errMb.Error(), http.StatusBadRequest)
		}
		json.NewEncoder(resWri).Encode(members)

	case http.MethodPut:
		var newMembererData model.TeamMember
		errDecode := json.NewDecoder(requ.Body).Decode(&newMembererData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
		}
		errAdd := db.AddMember(newMembererData)
		if errAdd != nil {
			http.Error(resWri, errAdd.Error(), http.StatusBadRequest)
		}
		resWri.Write([]byte("Member added"))

	case http.MethodPost:
		var updatedMembererData model.TeamMember
		errDecode := json.NewDecoder(requ.Body).Decode(&updatedMembererData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
		}
		errUpd := db.UpdateMember(updatedMembererData)
		if errUpd != nil {
			http.Error(resWri, errUpd.Error(), http.StatusBadRequest)
		}
		resWri.Write([]byte("Member updated"))

	case http.MethodDelete:
		var deletedMenberData model.TeamMember
		errDecode := json.NewDecoder(requ.Body).Decode(&deletedMenberData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
		}
		errDel := db.DeleteMember(deletedMenberData)
		if errDel != nil {
			http.Error(resWri, errDel.Error(), http.StatusBadRequest)
		}
		resWri.Write([]byte("Member deleted"))
	}
}
