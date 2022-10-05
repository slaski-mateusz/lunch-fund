package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/slaski-mateusz/lunch-fund/db"
	"github.com/slaski-mateusz/lunch-fund/model"
)

func membersHandler(resWri http.ResponseWriter, requ *http.Request) {
	switch requ.Method {
	case "GET":
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

	case "PUT":
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

	case "POST":
		var updateMembererData model.TeamMember
		errDecode := json.NewDecoder(requ.Body).Decode(&updateMembererData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
		}
		fmt.Printf("%+v\n\n", updateMembererData)
		// TODO Update mamber

	case "DELETE":
		var deleteMenberData model.TeamMember
		errDecode := json.NewDecoder(requ.Body).Decode(&deleteMenberData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
		}
		errDel := db.DeleteMember(deleteMenberData)
		if errDel != nil {
			http.Error(resWri, errDel.Error(), http.StatusBadRequest)
		}
		resWri.Write([]byte("Member deleted"))
	}
}
