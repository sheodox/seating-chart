package controllers

import (
	"database/sql"

	"github.com/sheodox/seating-chart/interactors"
)

type GuestController struct {
	Interactor interactors.GuestInteractor
}

type NewGuestParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	People    int    `json:"people"`
}

type GuestDTO struct {
	Id        string `json:"id"`
	TableId   string `json:"tableId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	People    int    `json:"people"`
}

func (g GuestController) List() ([]GuestDTO, error) {
	guestDto := make([]GuestDTO, 0)

	guestEntities, err := g.Interactor.List()

	if err != nil {
		return guestDto, err
	}

	for _, ent := range guestEntities {
		guest := GuestDTO{}
		guest.Id = ent.Id
		guest.LastName = ent.LastName
		guest.FirstName = ent.FirstName
		guest.People = ent.People
		guest.TableId = ent.TableId.String
		guestDto = append(guestDto, guest)
	}

	return guestDto, nil
}

func (g GuestController) Add(data map[string]any) (GuestDTO, error) {
	guestDto := GuestDTO{}
	guestReq := NewGuestParams{}

	// todo safety check
	guestReq.FirstName = data["firstName"].(string)
	guestReq.LastName = data["lastName"].(string)
	guestReq.People = int(data["people"].(float64))

	guestEntity, err := g.Interactor.Add(guestReq.FirstName, guestReq.LastName, guestReq.People)

	if err != nil {
		return guestDto, err
	}

	guestDto.Id = guestEntity.Id
	guestDto.LastName = guestEntity.LastName
	guestDto.FirstName = guestEntity.FirstName
	guestDto.People = guestEntity.People
	guestDto.TableId = guestEntity.TableId.String

	return guestDto, nil
}

func (g GuestController) Edit(data map[string]any) (GuestDTO, error) {
	guestDto := GuestDTO{}
	guestReq := GuestDTO{}

	// todo safety check
	guestReq.Id = data["id"].(string)
	guestReq.FirstName = data["firstName"].(string)
	guestReq.LastName = data["lastName"].(string)
	guestReq.People = int(data["people"].(float64))

	guestEntity, err := g.Interactor.Edit(guestReq.Id, guestReq.FirstName, guestReq.LastName, guestReq.People)

	if err != nil {
		return guestDto, err
	}

	guestDto.Id = guestEntity.Id
	guestDto.LastName = guestEntity.LastName
	guestDto.FirstName = guestEntity.FirstName
	guestDto.People = guestEntity.People
	guestDto.TableId = guestEntity.TableId.String

	return guestDto, nil
}

func (g GuestController) Assign(data map[string]any) (GuestDTO, error) {
	guestId := data["guestId"].(string)
	tableId, hasValidTableId := data["tableId"].(string)
	guestDto := GuestDTO{}

	guestEntity, err := g.Interactor.Assign(guestId, sql.NullString{String: tableId, Valid: hasValidTableId})

	if err != nil {
		return guestDto, err
	}

	guestDto.Id = guestEntity.Id
	guestDto.LastName = guestEntity.LastName
	guestDto.FirstName = guestEntity.FirstName
	guestDto.People = guestEntity.People
	guestDto.TableId = guestEntity.TableId.String

	return guestDto, nil
}

func (g GuestController) Delete(data map[string]any) (string, error) {
	id, ok := data["id"].(string)

	if !ok {
		return "", ErrBadData
	}

	err := g.Interactor.Delete(id)

	return id, err
}
