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
		var newOrderData model.Order
		errDecode := json.NewDecoder(requ.Body).Decode(&newOrderData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		}
		if newOrderData.TeamName == "" {
			http.Error(resWri, "Team name can not to be emty!", http.StatusBadRequest)
			return
		}
		errAdd := db.AddOrder(newOrderData)
		if errAdd != nil {
			http.Error(resWri, errAdd.Error(), http.StatusBadRequest)
			return
		}
		resWri.Write([]byte("Order added"))
		return

	case http.MethodPost:
		var updatedOrderData model.Order
		errDecode := json.NewDecoder(requ.Body).Decode(&updatedOrderData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		}
		if updatedOrderData.TeamName == "" {
			http.Error(resWri, "Team name can not to be emty!", http.StatusBadRequest)
			return
		}
		errUpd := db.UpdateOrder(updatedOrderData)
		if errUpd != nil {
			http.Error(resWri, errUpd.Error(), http.StatusBadRequest)
			return
		}
		resWri.Write([]byte("Order updated"))
		return

	case http.MethodDelete:
		var deletedOrderData model.Order
		errDecode := json.NewDecoder(requ.Body).Decode(&deletedOrderData)
		if errDecode != nil {
			http.Error(resWri, errDecode.Error(), http.StatusBadRequest)
			return
		}
		if deletedOrderData.TeamName == "" {
			http.Error(resWri, "Team name can not to be emty!", http.StatusBadRequest)
			return
		}
		errDel := db.DeleteOrder(deletedOrderData)
		if errDel != nil {
			http.Error(resWri, errDel.Error(), http.StatusBadRequest)
			return
		}
		resWri.Write([]byte("Order deleted"))
		return
	}
}
