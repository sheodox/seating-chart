package repositories

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/sheodox/seating-chart/entities"
)

type Line struct {
	db *sqlx.DB
}

type lineCoordinateSerialized struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type lineSerialized struct {
	Id     string `db:"id"`
	Coords string `db:"coords"`
	Closed bool   `db:"closed"`
}

func toLineEntity(l lineSerialized) entities.Line {
	coords := []entities.LineCoordinate{}
	json.Unmarshal([]byte(l.Coords), &coords)

	return entities.Line{
		Id:     l.Id,
		Closed: l.Closed,
		Coords: coords,
	}
}

func (l *Line) Add() (entities.Line, error) {
	id := genId()

	_, err := l.db.Exec("insert into lines (id, closed) values ($1, $2)", id, true)

	if err != nil {
		return entities.Line{}, err
	}

	return l.Get(id)
}

func (l *Line) Edit(ent entities.Line) (entities.Line, error) {
	coordsSerialized := []lineCoordinateSerialized{}

	for _, c := range ent.Coords {
		coordsSerialized = append(coordsSerialized, lineCoordinateSerialized{
			X: c.X,
			Y: c.Y,
		})
	}

	coordsMarshalled, err := json.Marshal(coordsSerialized)

	if err != nil {
		return entities.Line{}, err
	}

	_, err = l.db.Exec("update lines set closed=$1, coords=$2 where id=$3", ent.Closed, coordsMarshalled, ent.Id)

	if err != nil {
		return entities.Line{}, err
	}

	return l.Get(ent.Id)
}

func (l *Line) Get(id string) (entities.Line, error) {
	line := lineSerialized{}

	err := l.db.Get(&line, "select * from lines where id=$1", id)

	return toLineEntity(line), err
}

func (l *Line) List() ([]entities.Line, error) {
	lines := []entities.Line{}
	linesSerialized := []lineSerialized{}

	err := l.db.Select(&linesSerialized, "select * from lines")

	for _, l := range linesSerialized {
		lines = append(lines, toLineEntity(l))
	}

	return lines, err
}

func (l *Line) Delete(id string) error {
	_, err := l.db.Exec("delete from lines where id=$1", id)

	return err
}
