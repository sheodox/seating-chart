package repositories

import (
	"github.com/rs/xid"
	"github.com/sheodox/seating-chart/db"
)

type Repositories struct {
	Guest Guest
	Table Table
	Line  Line
}

func NewRepositories() *Repositories {
	return &Repositories{
		Guest: Guest{db.Connection},
		Table: Table{db.Connection},
		Line:  Line{db.Connection},
	}
}

func genId() string {
	return xid.New().String()
}
