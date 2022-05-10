package repositories

import (
	"github.com/rs/xid"
	"github.com/sheodox/seating-chart/db"
)

type Repositories struct {
	Guest Guest
	Table Table
}

func NewRepositories() *Repositories {
	return &Repositories{
		Guest: Guest{db.Connection},
		Table: Table{db.Connection},
	}
}

func genId() string {
	return xid.New().String()
}
