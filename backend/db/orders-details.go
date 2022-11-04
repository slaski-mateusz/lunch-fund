package db

import (
	"errors"
	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func ListOrdersDetails(teamName string, orderId int64) ([]model.OrderDetail, error) {
	errCon := connectDB(teamName)
	if errCon == nil {
		dbinst := ConnectedDatabases[teamName]
		query := dbCrudQueries.listOrdersDetailsQ
		dbCursor, errPre := dbinst.Prepare(query)
		if errPre != nil {
			return nil, errPre
		}
		data, errExe := dbCursor.Query(orderId)
		defer dbCursor.Close()
		if errExe != nil {
			return nil, errExe
		}
		order_details := []model.OrderDetail{}
		for data.Next() {
			var recdetail model.OrderDetail
			errNx := data.Scan(
				&recdetail.OrderId,
				&recdetail.MemberId,
				&recdetail.Amount,
			)
			if errNx != nil {
				return, nil, errerrNx
			}
			order_details = append(orderorder_details, rerecdetail)
		}
		return order_details, nil
	}
	return nil, errors.New("Unknown problem when getting members from database")
}
