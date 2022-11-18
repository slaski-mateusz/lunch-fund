package db

import (
	"database/sql"
	"errors"

	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func ListOrdersDetails(teamName string, orderId int64) ([]model.OrderDetail, error) {
	errCon := connectDB(teamName)
	if errCon != nil {
		return nil, errCon
	} else {
		dbinst := ConnectedDatabases[teamName]
		dbCursor, errPre := dbinst.Prepare(
			dbCrudQueries.ordersDetailsListQ,
		)
		if errPre != nil {
			return nil, errPre
		}
		data, errExe := dbCursor.Query(orderId)
		defer dbCursor.Close()
		if errExe != nil {
			return nil, errExe
		}
		ordersDetails := []model.OrderDetail{}
		for data.Next() {
			var recDetail model.OrderDetail
			errNx := data.Scan(
				&recDetail.OrderId,
				&recDetail.MemberId,
				&recDetail.IsFounder,
				&recDetail.Amount,
			)
			if errNx != nil {
				return nil, errNx
			}
			ordersDetails = append(ordersDetails, recDetail)
		}
		return ordersDetails, nil
	}
}

func AddOrderDetail(newOrderDetail model.OrderDetail) error {
	errCon := connectDB(newOrderDetail.TeamName)
	if errCon != nil {
		return errCon
	} else {
		dbinst := ConnectedDatabases[newOrderDetail.TeamName]
		_, errExe := dbinst.Exec(
			dbCrudQueries.orderDetailsAddQ,
			newOrderDetail.OrderId,
			newOrderDetail.MemberId,
			newOrderDetail.IsFounder,
			newOrderDetail.Amount,
		)
		if errExe != nil {
			return errExe
		}
		return nil
	}
}

func UpdateOrderDetail(orderDetailData model.OrderDetail) error {
	errCon := connectDB(orderDetailData.TeamName)
	if errCon != nil {
		return errCon
	} else {
		dbinst := ConnectedDatabases[orderDetailData.TeamName]
		_, errExe := dbinst.Exec(
			dbCrudQueries.orderDetailsUpdateQ,
			orderDetailData.OrderId,
			orderDetailData.MemberId,
			orderDetailData.IsFounder,
			orderDetailData.Amount,
			orderDetailData.OrderId,
			orderDetailData.MemberId,
		)
		if errExe != nil {
			return errExe
		}
		return nil
	}
}

func DeleteOrderDetail(orderDetailToDelete model.OrderDetail) error {
	errCon := connectDB(orderDetailToDelete.TeamName)
	if errCon != nil {
		return errCon
	} else {
		dbinst := ConnectedDatabases[orderDetailToDelete.TeamName]
		row := dbinst.QueryRow(
			dbCrudQueries.orderDetailCheckIfExistQ,
			orderDetailToDelete.OrderId,
			orderDetailToDelete.MemberId,
		)
		errQuer := row.Scan(
			&orderDetailToDelete.OrderId,
			&orderDetailToDelete.MemberId,
		)
		if errQuer != nil {
			if errQuer == sql.ErrNoRows {
				return errors.New("No such order detail in database")
			} else {
				return errQuer
			}
		} else {
			_, errExe := dbinst.Exec(
				dbCrudQueries.orderDetailDeleteQ,
				orderDetailToDelete.OrderId,
				orderDetailToDelete.MemberId,
			)
			if errExe != nil {
				return errExe
			}
			return nil
		}
	}
}
