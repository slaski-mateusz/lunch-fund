package db

var dbInitQueries = map[string]string{
	"activateForeginKeys": "PRAGMA foreign_keys = ON;",
	"createTableMembers": `CREATE TABLE IF NOT EXISTS members (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		phone TEXT NOT NULL UNIQUE,
		is_admin INTEGER NOT NULL DEFAULT 0,
		active INTEGER NOT NULL DEFAULT 1,
		avatar BLOB)`,
	"createTableOrders": `CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		timestamp INTEGER NOT NULL,
		founder_id INTEGER NOT NULL,
		delivery_cost INTEGER DEFAULT 0,
		tip_cost  INTEGER DEFAULT 0,
		FOREIGN KEY (founder_id) REFERENCES members (id) ON UPDATE CASCADE ON DELETE RESTRICT)`,
	"createTableOrdersDetails": `CREATE TABLE IF NOT EXISTS orders_details (
		order_id INTEGER NOT NULL,
		member_id INTEGER NOT NULL,
		amount INTEGER NOT NULL,
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

var dbModQueries = struct {
	add1stAdminQ        string
	addMemberQ          string
	updateMemberQ       string
	deleteMemberQ       string
	listMembersQ        string
	addOrderQ           string
	updateOrderQ        string
	deleteOrderQ        string
	listOrdersQ         string
	addOrderDetailsQ    string
	updateOrderDetailsQ string
	deleteOrderDetailsQ string
	listOrdersDetailsQ  string
}{
	add1stAdminQ:        ``,
	addMemberQ:          ``,
	updateMemberQ:       ``,
	deleteMemberQ:       ``,
	listMembersQ:        `SELECT * FROM members;`,
	addOrderQ:           ``,
	updateOrderQ:        ``,
	deleteOrderQ:        ``,
	listOrdersQ:         `SELECT * FROM orders;`,
	addOrderDetailsQ:    ``,
	updateOrderDetailsQ: ``,
	deleteOrderDetailsQ: ``,
	listOrdersDetailsQ:  `SELECT * FROM orders_details;`,
}
