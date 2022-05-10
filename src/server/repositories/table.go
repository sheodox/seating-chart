package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/sheodox/seating-chart/entities"
)

type Table struct {
	db *sqlx.DB
}

func (g *Table) Add(name string, posX, posY float64) (entities.Table, error) {
	id := genId()

	_, err := g.db.Exec("insert into tables (name, pos_x, pos_y, id, capacity) values ($1, $2, $3, $4, $5)", name, posX, posY, id, 8)

	if err != nil {
		return entities.Table{}, err
	}

	return g.Get(id)
}

func (g *Table) Edit(id, name string, posX, posY float64, capacity int) (entities.Table, error) {
	_, err := g.db.Exec("update tables set name=$2, pos_x=$3, pos_y=$4, capacity=$5 where id=$1", id, name, posX, posY, capacity)

	if err != nil {
		return entities.Table{}, err
	}

	return g.Get(id)
}

func (g *Table) Get(id string) (entities.Table, error) {
	table := entities.Table{}

	err := g.db.Get(&table, "select * from tables where id=$1", id)

	return table, err
}

func (g *Table) List() ([]entities.Table, error) {
	tables := []entities.Table{}

	err := g.db.Select(&tables, "select * from tables")

	return tables, err
}

func (g *Table) Delete(id string) error {
	_, err := g.db.Exec("delete from tables where id=$1", id)

	return err
}

//-- name: GetGuests :many
//select * from guests

//-- name: CreateGuest :one
//insert into guests (
//id, first_name, last_name, people
//) values ($1, $2, $3, $4)

//-- name: DeleteGuest :exec
//delete from guests
//where id=$1
