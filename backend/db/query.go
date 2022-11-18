package db

var activateForeginKeys = "PRAGMA foreign_keys = ON;"

var dbInitQueries = map[string]string{
	"createTableMembers": `
	    CREATE TABLE IF NOT EXISTS members (
		id INTEGER PRIMARY KEY,
		member_name TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		phone TEXT NOT NULL UNIQUE,
		is_admin INTEGER NOT NULL DEFAULT 0,
		is_active INTEGER NOT NULL DEFAULT 1,
		secret TEXT)`,
	"createTableOrders": `CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY,
		order_name TEXT NOT NULL UNIQUE,
		timestamp INTEGER NOT NULL,
		delivery_cost INTEGER DEFAULT 0,
		tip_cost  INTEGER DEFAULT 0)`,
	"createTableOrdersDetails": `CREATE TABLE IF NOT EXISTS orders_details (
		order_id INTEGER NOT NULL,
		member_id INTEGER NOT NULL,
		is_founder INTEGER NOT NULL,
		amount INTEGER NOT NULL,
		PRIMARY KEY (order_id, member_id),
		FOREIGN KEY (order_id) REFERENCES orders (id) ON UPDATE CASCADE ON DELETE RESTRICT,
		FOREIGN KEY (member_id) REFERENCES members (id) ON UPDATE CASCADE ON DELETE RESTRICT)`,
	"createTableDebts": `CREATE TABLE IF NOT EXISTS debts (
		debtor_id INTEGER NOT NULL,
		creditor_id INTEGER NOT NULL,
		amount INTEGER NOT NULL,
		return_timestamp INTEGER,
		FOREIGN KEY (debtor_id) REFERENCES members (id) ON UPDATE CASCADE ON DELETE RESTRICT,
		FOREIGN KEY (creditor_id) REFERENCES members (id) ON UPDATE CASCADE ON DELETE RESTRICT)`,
}

var dbCrudQueries = struct {
	memberAddQ               string
	memberCheckIfExistQ      string
	memberUpdateQ            string
	memberDeleteMQ           string
	membersListQ             string
	orderAddQ                string
	orderCheckIfExistQ       string
	orderUpdateQ             string
	orderDeleteQ             string
	ordersListQ              string
	orderDetailsAddQ         string
	orderDetailsUpdateQ      string
	orderDetailCheckIfExistQ string
	orderDetailDeleteQ       string
	ordersDetailsListQ       string
}{
	memberAddQ:               `INSERT INTO members (member_name, email, phone, is_admin, is_active) VALUES (?, ?, ?, ?, ?);`,
	memberCheckIfExistQ:      `SELECT id FROM members WHERE id=?`,
	memberUpdateQ:            `UPDATE members SET member_name=?, email=?, phone=?, is_admin=?, is_active=? WHERE id=?`,
	memberDeleteMQ:           `DELETE FROM members WHERE id=?`,
	membersListQ:             `SELECT * FROM members;`,
	orderAddQ:                `INSERT INTO orders (order_name, timestamp, delivery_cost, tip_cost) VALUES (?, ?, ?, ?);`,
	orderCheckIfExistQ:       `SELECT id FROM orders WHERE id=?`,
	orderUpdateQ:             `UPDATE orders SET order_name=?, timestamp=?, delivery_cost=?, tip_cost=? WHERE id=?`,
	orderDeleteQ:             `DELETE FROM orders WHERE id=?`,
	ordersListQ:              `SELECT * FROM orders;`,
	orderDetailsAddQ:         `INSERT INTO orders_details (order_id, member_id, is_founder, amount) VALUES (?, ?, ?, ?);`,
	orderDetailsUpdateQ:      `UPDATE orders_details SET order_id=?, member_id=?, is_founder=?, amount=? WHERE order_id=? AND member_id=?`,
	orderDetailCheckIfExistQ: `SELECT order_id, member_id FROM orders_details WHERE order_id=? AND member_id=?`,
	orderDetailDeleteQ:       `DELETE FROM orders_details WHERE order_id=? AND member_id=?`,
	ordersDetailsListQ:       `SELECT * FROM orders_details WHERE order_id=?;`,
}
