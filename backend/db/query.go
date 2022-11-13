package db

var dbInitQueries = map[string]string{
	"activateForeginKeys": "PRAGMA foreign_keys = ON;",
	"createTableMembers": `CREATE TABLE IF NOT EXISTS members (
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
	addMemberQ          string
	checkIfMemberExistQ string
	updateMemberQ       string
	deleteMemberQ       string
	listMembersQ        string
	addOrderQ           string
	checkIfOrderExistQ  string
	updateOrderQ        string
	deleteOrderQ        string
	listOrdersQ         string
	addOrderDetailsQ    string
	updateOrderDetailsQ string
	deleteOrderDetailsQ string
	listOrdersDetailsQ  string
}{
	addMemberQ:          `INSERT INTO members (member_name, email, phone, is_admin, is_active) VALUES (?, ?, ?, ?, ?);`,
	checkIfMemberExistQ: `SELECT id FROM members WHERE id=?`,
	updateMemberQ:       `UPDATE members SET member_name=?, email=?, phone=?, is_admin=?, is_active=? WHERE id=?`,
	deleteMemberQ:       `DELETE FROM members WHERE id=?`,
	listMembersQ:        `SELECT * FROM members;`,
	addOrderQ:           `INSERT INTO orders (order_name, timestamp, founder_id, delivery_cost, tip_cost) VALUES (?, ?, ?, ?, ?);`,
	checkIfOrderExistQ:  `SELECT id FROM orders WHERE id=?`,
	updateOrderQ:        `UPDATE orders SET order_name=?, timestamp=?, founder_id=?, delivery_cost=?, tip_cost=? WHERE id=?`,
	deleteOrderQ:        `DELETE FROM orders WHERE id=?`,
	listOrdersQ:         `SELECT * FROM orders;`,
	addOrderDetailsQ:    `INSERT INTO orders_details (order_id, member_id, is_founder, amount) VALUES (?, ?, ?, ?);`,
	updateOrderDetailsQ: `UPDATE orders_details SET order_id=?, member_id=?, is_founder=?, amount=? WHERE order_id=? AND member_id=?`,
	deleteOrderDetailsQ: `DELETE FROM orders_details WHERE order_id=? AND member_id=?`,
	listOrdersDetailsQ:  `SELECT * FROM orders_details WHERE order_id=?;`,
}
