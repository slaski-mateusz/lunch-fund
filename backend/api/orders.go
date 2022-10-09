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
		}
		orders, errMb := db.ListOrders(teamData.TeamName)
		if errMb != nil {
			http.Error(resWri, errMb.Error(), http.StatusBadRequest)
		}
		json.NewEncoder(resWri).Encode(orders)

	case http.MethodPut:
		json.NewEncoder(resWri).Encode("Not implemented yet")
	case http.MethodPost:
		json.NewEncoder(resWri).Encode("Not implemented yet")
	case http.MethodDelete:
		json.NewEncoder(resWri).Encode("Not implemented yet")
	}
}
