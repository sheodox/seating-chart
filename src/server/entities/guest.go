package entities

import "database/sql"

type Guest struct {
	Id        string         `db:"id"`
	TableId   sql.NullString `db:"table_id"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	People    int            `db:"people"`
	Notes     string         `db:"notes"`
	Going     bool           `db:"going"`
}
