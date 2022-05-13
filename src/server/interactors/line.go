package interactors

import (
	"github.com/sheodox/seating-chart/entities"
	"github.com/sheodox/seating-chart/repositories"
)

type LineInteractor struct {
	Repo repositories.Line
}

func (l *LineInteractor) Add() (entities.Line, error) {
	return l.Repo.Add()
}

func (l *LineInteractor) Edit(ent entities.Line) (entities.Line, error) {
	for i, coord := range ent.Coords {
		ent.Coords[i].X = clampPos(coord.X, 0.02)
		ent.Coords[i].Y = clampPos(coord.Y, 0.02)
	}
	return l.Repo.Edit(ent)
}

func (l *LineInteractor) Delete(id string) error {
	if id == "" {
		return ErrInvalidId
	}

	return l.Repo.Delete(id)
}

func (l *LineInteractor) List() ([]entities.Line, error) {
	return l.Repo.List()
}
