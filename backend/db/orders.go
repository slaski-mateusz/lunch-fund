package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func ListOrders(teamName string) ([]model.Order, error) {
	errCon := connectDB(teamName)
	if errCon != nil {
		return nil, errCon
	} else {
		dbinst := ConnectedDatabases[teamName]
		query := dbCrudQueries.ordersListQ
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
}

func AddOrder(newOrder model.Order) error {
	errCon := connectDB(newOrder.TeamName)
	if errCon != nil {
		return errCon
	} else {
		dbinst := ConnectedDatabases[newOrder.TeamName]
		_, errExe := dbinst.Exec(
			dbCrudQueries.orderAddQ,
			newOrder.OrderName,
			newOrder.Timestamp,
			newOrder.DeliveryCost,
			newOrder.TipCost,
		)
		if errExe != nil {
			return errExe
		}
		return nil
	}
}

func UpdateOrder(orderData model.Order) error {
	errCon := connectDB((orderData.TeamName))
	if errCon != nil {
		return errCon
	} else {
		dbinst := ConnectedDatabases[orderData.TeamName]
		_, errExe := dbinst.Exec(
			dbCrudQueries.orderUpdateQ,
			orderData.OrderName,
			orderData.Timestamp,
			orderData.DeliveryCost,
			orderData.TipCost,
			orderData.Id,
		)
		if errExe != nil {
			return errExe
		}
		return nil
	}
}

func DeleteOrder(deletedOrder model.Order) error {
	errCon := connectDB(deletedOrder.TeamName)
	if errCon != nil {
		return errCon
	} else {
		dbinst := ConnectedDatabases[deletedOrder.TeamName]
		row := dbinst.QueryRow(
			dbCrudQueries.orderCheckIfExistQ,
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
				dbCrudQueries.orderDeleteQ,
				deletedOrder.Id,
			)
			if errExe != nil {
				return errExe
			}
			return nil
		}
	}
}
