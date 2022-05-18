package controllers

import (
	"database/sql"

	"github.com/sheodox/seating-chart/entities"
	"github.com/sheodox/seating-chart/interactors"
)

type GuestController struct {
	Interactor interactors.GuestInteractor
}

type NewGuestParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	People    int    `json:"people"`
	Going     bool   `json:"going"`
}

type GuestDTO struct {
	Id        string `json:"id"`
	TableId   string `json:"tableId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	People    int    `json:"people"`
	Going     bool   `json:"going"`
}

func guestEntityToDto(ent entities.Guest) GuestDTO {
	return GuestDTO{
		Id:        ent.Id,
		LastName:  ent.LastName,
		FirstName: ent.FirstName,
		People:    ent.People,
		TableId:   ent.TableId.String,
		Going:     ent.Going,
	}
}

func (g GuestController) List() ([]GuestDTO, error) {
	guestDto := make([]GuestDTO, 0)

	guestEntities, err := g.Interactor.List()

	if err != nil {
		return guestDto, err
	}

	for _, ent := range guestEntities {
		guestDto = append(guestDto, guestEntityToDto(ent))
	}

	return guestDto, nil
}

func (g GuestController) Add(data map[string]any) (GuestDTO, error) {
	guestReq := NewGuestParams{}

	// todo safety check
	guestReq.FirstName = data["firstName"].(string)
	guestReq.LastName = data["lastName"].(string)
	guestReq.People = int(data["people"].(float64))
	guestReq.Going = data["going"].(bool)

	guestEntity, err := g.Interactor.Add(guestReq.FirstName, guestReq.LastName, guestReq.People, guestReq.Going)

	if err != nil {
		return GuestDTO{}, err
	}

	return guestEntityToDto(guestEntity), nil

}

func (g GuestController) Edit(data map[string]any) (GuestDTO, error) {
	guestReq := GuestDTO{}

	// todo safety check
	guestReq.Id = data["id"].(string)
	guestReq.FirstName = data["firstName"].(string)
	guestReq.LastName = data["lastName"].(string)
	guestReq.People = int(data["people"].(float64))
	guestReq.Going = data["going"].(bool)

	guestEntity, err := g.Interactor.Edit(guestReq.Id, guestReq.FirstName, guestReq.LastName, guestReq.People, guestReq.Going)

	if err != nil {
		return GuestDTO{}, err
	}

	return guestEntityToDto(guestEntity), nil
}

func (g GuestController) Assign(data map[string]any) (GuestDTO, error) {
	guestId := data["guestId"].(string)
	tableId, hasValidTableId := data["tableId"].(string)

	guestEntity, err := g.Interactor.Assign(guestId, sql.NullString{String: tableId, Valid: hasValidTableId})

	if err != nil {
		return GuestDTO{}, err
	}

	return guestEntityToDto(guestEntity), nil
}

func (g GuestController) Delete(data map[string]any) (string, error) {
	id, ok := data["id"].(string)

	if !ok {
		return "", ErrBadData
	}

	err := g.Interactor.Delete(id)

	return id, err
}
