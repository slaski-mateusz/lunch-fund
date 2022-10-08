package db

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/slaski-mateusz/lunch-fund/backend/model"
)

func ListOrders(team string) error {
	// TODO: Listing orders
	return errors.New("Unknown problem when listing orders")
}

func AddOrder(team string, newOrder model.Order) error {
	// TODO: adding new order to database
	return errors.New("Unknown problem when adding order to database")
}

func UpdateOrder(team string, newOrder model.Order) error {
	// TODO: adding new order to database
	return errors.New("Unknown problem when updating order in database")
}

func DeleteOrder(team string, newOrder model.Order) error {
	// TODO: adding new order to database
	return errors.New("Unknown problem when removing order from database")
}
