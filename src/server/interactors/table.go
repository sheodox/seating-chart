package interactors

import (
	"github.com/sheodox/seating-chart/entities"
	"github.com/sheodox/seating-chart/repositories"
)

type TableInteractor struct {
	Repo repositories.Table
}

func (g *TableInteractor) validateTable(name string) error {
	if name == "" {
		return ErrInvalidTableName
	}

	return nil
}

func (g *TableInteractor) Add(name string, posX, posY float64) (entities.Table, error) {
	err := g.validateTable(name)

	if err != nil {
		return entities.Table{}, err
	}

	return g.Repo.Add(name, posX, posY)
}

func (g *TableInteractor) Edit(id, name string, posX, posY float64, capacity int) (entities.Table, error) {
	err := g.validateTable(name)

	if err != nil {
		return entities.Table{}, err
	}

	return g.Repo.Edit(id, name, posX, posY, capacity)
}

func (g *TableInteractor) Delete(id string) error {
	if id == "" {
		return ErrInvalidId
	}

	return g.Repo.Delete(id)
}

func (g *TableInteractor) List() ([]entities.Table, error) {
	return g.Repo.List()
}
