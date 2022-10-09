package db

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func ListOrders(team string) ([]model.Order, error) {
	// TODO: Listing orders
	return nil, errors.New("Listing orders not implemented in database")
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
