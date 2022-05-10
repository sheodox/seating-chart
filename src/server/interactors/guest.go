package interactors

import (
	"database/sql"

	"github.com/sheodox/seating-chart/entities"
	"github.com/sheodox/seating-chart/repositories"
)

type GuestInteractor struct {
	Repo repositories.Guest
}

func (g *GuestInteractor) validateGuest(firstName, lastName string, people int) error {
	if people < 1 {
		return ErrInvalidPeople
	}

	if firstName == "" || lastName == "" {
		return ErrInvalidName
	}

	return nil
}

func (g *GuestInteractor) Add(firstName, lastName string, people int) (entities.Guest, error) {
	err := g.validateGuest(firstName, lastName, people)

	if err != nil {
		return entities.Guest{}, err
	}

	return g.Repo.Add(firstName, lastName, people)
}

func (g *GuestInteractor) Edit(id, firstName, lastName string, people int) (entities.Guest, error) {
	err := g.validateGuest(firstName, lastName, people)

	if err != nil {
		return entities.Guest{}, err
	}

	return g.Repo.Edit(id, firstName, lastName, people)
}

func (g *GuestInteractor) Assign(guestId string, tableId sql.NullString) (entities.Guest, error) {
	return g.Repo.Assign(guestId, tableId)
}

func (g *GuestInteractor) Delete(id string) error {
	if id == "" {
		return ErrInvalidId
	}

	return g.Repo.Delete(id)
}

func (g *GuestInteractor) List() ([]entities.Guest, error) {
	return g.Repo.List()
}
