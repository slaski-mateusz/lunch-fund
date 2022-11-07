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
		ordersDetails := []model.OrderDetail{}
		for data.Next() {
			var recDetail model.OrderDetail
			errNx := data.Scan(
				&recDetail.OrderId,
				&recDetail.MemberId,
				&recDetail.Amount,
			)
			if errNx != nil {
				return nil, errNx
			}
			ordersDetails = append(ordersDetails, recDetail)
		}
		return ordersDetails, nil
	}
	return nil, errors.New("Unknown problem when getting members from database")
}
