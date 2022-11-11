package api

import (
	"encoding/json"
	"net/http"

	"github.com/slaski-mateusz/lunch-fund/backend/db"
	"github.com/slaski-mateusz/lunch-fund/backend/model"
	// "github.com/slaski-mateusz/lunch-fund/backend/model"
)

func ordersDetailsHandler(resWri http.ResponseWriter, requ *http.Request) {
	switch requ.Method {
	case http.MethodGet:
		var orderDetailsRequest model.Order
		errDecode := json.NewDecoder(requ.Body).Decode(&orderDetailsRequest)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
		}
		if orderDetailsRequest.TeamName == "" {
			http.Error(resWri, "API says: Team name can not to be emty!", http.StatusBadRequest)
			return
		}
		orders, errOrd := db.ListOrdersDetails(
			orderDetailsRequest.TeamName,
			orderDetailsRequest.Id,
		)
		if errOrd != nil {
			http.Error(resWri, errOrd.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(resWri).Encode(orders)
		return
	case http.MethodPut:

	case http.MethodPost:

	case http.MethodDelete:

	}
}
