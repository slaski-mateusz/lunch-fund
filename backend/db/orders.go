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

func AddOrder(newOrder model.Order) error {
	errCon := connectDB(newOrder.TeamName)
	if errCon == nil {
		dbinst := ConnectedDatabases[newOrder.TeamName]
		_, errExe := dbinst.Exec(
			dbCrudQueries.addOrderQ,
			newOrder.OrderName,
			newOrder.Timestamp,
			newOrder.FounderId,
			newOrder.DeliveryCost,
			newOrder.TipCost,
		)
		if errExe != nil {
			return errExe
		}
		return nil
	}
	return errors.New("Unknown problem when adding order to database")
}

func UpdateOrder(orderData model.Order) error {
	errCon := connectDB((orderData.TeamName))
	if errCon == nil {
		dbinst := ConnectedDatabases[orderData.TeamName]
		_, errExe := dbinst.Exec(
			dbCrudQueries.updateOrderQ,
			orderData.OrderName,
			orderData.Timestamp,
			orderData.FounderId,
			orderData.DeliveryCost,
			orderData.TipCost,
			orderData.Id,
		)
		if errExe != nil {
			return errExe
		}
		return nil
	}
	return errors.New("Unknown problem when updating member in database")
}

func DeleteOrder(deletedOrder model.Order) error {
	errCon := connectDB(deletedOrder.TeamName)
	if errCon == nil {
		dbinst := ConnectedDatabases[deletedOrder.TeamName]
		row := dbinst.QueryRow(
			dbCrudQueries.checkIfOrderExistQ,
			deletedOrder.Id,
		)
		errQuer := row.Scan(&deletedOrder.Id)
		if errQuer != nil {
			if errQuer == sql.ErrNoRows {
				return errors.New("No such user in database")
			} else {
				return errQuer
			}
		} else {
			_, errExe := dbinst.Exec(
				dbCrudQueries.deleteOrderQ,
				deletedOrder.Id,
			)
			if errExe != nil {
				return errExe
			}
			return nil
		}
	}
	return errors.New("Unknown problem when delete member from database")
}
