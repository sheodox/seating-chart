package controllers

import (
	"github.com/sheodox/seating-chart/interactors"
)

type TableController struct {
	Interactor interactors.TableInteractor
}

type NewTableParams struct {
	Name string  `json:"name"`
	PosX float64 `json:"posX"`
	PosY float64 `json:"posY"`
}

type TableDTO struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	PosX     float64 `json:"posX"`
	PosY     float64 `json:"posY"`
	Capacity int     `json:"capacity"`
}

func (g TableController) List() ([]TableDTO, error) {
	tableDto := make([]TableDTO, 0)

	guestEntities, err := g.Interactor.List()

	if err != nil {
		return tableDto, err
	}

	for _, ent := range guestEntities {
		guest := TableDTO{}
		guest.Id = ent.Id
		guest.Name = ent.Name
		guest.Capacity = ent.Capacity
		guest.PosX = ent.PosX
		guest.PosY = ent.PosY
		tableDto = append(tableDto, guest)
	}

	return tableDto, nil
}

func (g TableController) Add(data map[string]any) (TableDTO, error) {
	tableDto := TableDTO{}
	tableReq := NewTableParams{}

	// todo safety check
	tableReq.Name = data["name"].(string)
	tableReq.PosX = data["posX"].(float64)
	tableReq.PosY = data["posY"].(float64)

	tableEntity, err := g.Interactor.Add(tableReq.Name, tableReq.PosX, tableReq.PosY)

	if err != nil {
		return tableDto, err
	}

	tableDto.Id = tableEntity.Id
	tableDto.Name = tableEntity.Name
	tableDto.PosX = tableEntity.PosX
	tableDto.PosY = tableEntity.PosY
	tableDto.Capacity = tableEntity.Capacity

	return tableDto, nil
}

func (g TableController) Edit(data map[string]any) (TableDTO, error) {
	tableDto := TableDTO{}
	tableReq := TableDTO{}

	// todo safety check
	tableReq.Id = data["id"].(string)
	tableReq.Name = data["name"].(string)
	tableReq.PosX = data["posX"].(float64)
	tableReq.PosY = data["posY"].(float64)
	tableReq.Capacity = int(data["capacity"].(float64))

	tableEntity, err := g.Interactor.Edit(tableReq.Id, tableReq.Name, tableReq.PosX, tableReq.PosY, tableReq.Capacity)

	if err != nil {
		return tableDto, err
	}

	tableDto.Id = tableEntity.Id
	tableDto.Name = tableEntity.Name
	tableDto.PosX = tableEntity.PosX
	tableDto.PosY = tableEntity.PosY
	tableDto.Capacity = tableEntity.Capacity

	return tableDto, nil
}

func (g TableController) Delete(data map[string]any) (string, error) {
	id, ok := data["id"].(string)

	if !ok {
		return "", ErrBadData
	}

	err := g.Interactor.Delete(id)

	return id, err
}
