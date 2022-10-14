package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func ListOrders(teamName string) ([]model.Order, error) {
	errCon := connectDB(teamName)
	if errCon == nil {
		dbinst := ConnectedDatabases[teamName]
		query := dbCrudQueries.listOrdersQ
		dbCursor, errPre := dbinst.Prepare(query)
		if errPre != nil {
			return nil, errPre
		}
		data, errExe := dbCursor.Query()
		defer dbCursor.Close()
		if errExe != nil {
			return nil, errExe
		}
		orders := []model.Order{}
		for data.Next() {
			var recorder model.Order
			var deliveryCost sql.NullInt64
			var tipCost sql.NullInt64
			errNx := data.Scan(
				&recorder.Id,
				&recorder.OrderName,
				&recorder.Timestamp,
				&recorder.FounderId,
				&recorder.DeliveryCost,
				&recorder.TipCost,
			)
			recorder.DeliveryCost = model.Money(deliveryCost.Int64)
			recorder.TipCost = model.Money(tipCost.Int64)
			if errNx != nil {
				return nil, errNx
			}
			orders = append(orders, recorder)
		}
		return orders, nil
	}
	return nil, errors.New("Unknown problem when getting members from database")
}

func AddOrder(team string, newOrder model.Order) error {
	// TODO: adding new order to database
	return errors.New("Adding orders not implemented in database")
}

func UpdateOrder(team string, newOrder model.Order) error {
	// TODO: adding new order to database
	return errors.New("Updating orders not implemented in database")
}

func DeleteOrder(team string, newOrder model.Order) error {
	// TODO: adding new order to database
	return errors.New("Deleteing orders not implemented in database")
}
