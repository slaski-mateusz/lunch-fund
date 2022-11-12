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
		// Listing ordrt details for selected order id
		var ordersDetailsRequest model.Order
		errDecode := json.NewDecoder(requ.Body).Decode(&ordersDetailsRequest)
		if errDecode != nil {
			http.Error(
				resWri,
				errDecode.Error(),
				http.StatusBadRequest,
			)
			return
		}
		if ordersDetailsRequest.TeamName == "" {
			http.Error(
				resWri,
				"API says: Team name can not to be emty!",
				http.StatusBadRequest,
			)
			return
		}
		orders, errOrd := db.ListOrdersDetails(
			ordersDetailsRequest.TeamName,
			ordersDetailsRequest.Id,
		)
		if errOrd != nil {
			http.Error(
				resWri,
				errOrd.Error(),
				http.StatusBadRequest,
			)
			return
		}
		json.NewEncoder(resWri).Encode(orders)
		return
	case http.MethodPut:
		// Adding new single order detail
		var newOrderDetail model.OrderDetail
		errDecode := json.NewDecoder(requ.Body).Decode(&newOrderDetail)
		if errDecode != nil {
			http.Error(
				resWri,
				errDecode.Error(),
				http.StatusBadRequest,
			)
			return
		}
		if newOrderDetail.TeamName == "" {
			http.Error(
				resWri,
				"Team name can not to be emty!",
				http.StatusBadRequest,
			)
			return
		}
		errAdd := db.AddOrderDetail(newOrderDetail)
		if errAdd != nil {
			http.Error(
				resWri,
				errAdd.Error(),
				http.StatusBadRequest,
			)
			return
		}
		resWri.Write([]byte("Order detail added"))
		return
	case http.MethodPost:

	case http.MethodDelete:

	}
}
