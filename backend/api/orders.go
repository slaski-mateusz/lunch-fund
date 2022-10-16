package api

import (
	"encoding/json"
	"net/http"

	"github.com/slaski-mateusz/lunch-fund/backend/db"
	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func ordersHandler(resWri http.ResponseWriter, requ *http.Request) {
	switch requ.Method {
	case http.MethodGet:
		var teamData model.Team
		errDecode := json.NewDecoder(requ.Body).Decode(&teamData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		}
		if teamData.TeamName == "" {
			http.Error(resWri, "Team name can not to be emty!", http.StatusBadRequest)
			return
		}
		orders, errOrd := db.ListOrders(teamData.TeamName)
		if errOrd != nil {
			http.Error(resWri, errOrd.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(resWri).Encode(orders)
		return

	case http.MethodPut:
		json.NewEncoder(resWri).Encode("Not implemented in api yet")
		return
	case http.MethodPost:
		json.NewEncoder(resWri).Encode("Not implemented in api yet")
		return
	case http.MethodDelete:
		json.NewEncoder(resWri).Encode("Not implemented in api yet")
		return
	}
}
