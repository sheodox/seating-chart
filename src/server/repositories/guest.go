package repositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sheodox/seating-chart/entities"
)

type Guest struct {
	db *sqlx.DB
}

func (g *Guest) Add(firstName, lastName string, people int) (entities.Guest, error) {
	id := genId()

	_, err := g.db.Exec("insert into guests (first_name, last_name, people, id) values ($1, $2, $3, $4)", firstName, lastName, people, id)

	if err != nil {
		return entities.Guest{}, err
	}

	return g.Get(id)
}

func (g *Guest) Edit(id, firstName, lastName string, people int) (entities.Guest, error) {
	_, err := g.db.Exec("update guests set first_name=$2, last_name=$3, people=$4 where id=$1", id, firstName, lastName, people)

	if err != nil {
		return entities.Guest{}, err
	}

	return g.Get(id)
}

func (g *Guest) Assign(guestId string, tableId sql.NullString) (entities.Guest, error) {
	_, err := g.db.Exec("update guests set table_id=$2 where id=$1", guestId, tableId)

	if err != nil {
		return entities.Guest{}, err
	}

	return g.Get(guestId)
}

func (g *Guest) Get(id string) (entities.Guest, error) {
	guest := entities.Guest{}

	err := g.db.Get(&guest, "select * from guests where id=$1", id)

	return guest, err
}

func (g *Guest) List() ([]entities.Guest, error) {
	guests := []entities.Guest{}

	err := g.db.Select(&guests, "select * from guests")

	return guests, err
}

func (g *Guest) Delete(id string) error {
	_, err := g.db.Exec("delete from guests where id=$1", id)

	return err
}
