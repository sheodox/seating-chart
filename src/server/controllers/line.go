package controllers

import (
	"github.com/sheodox/seating-chart/entities"
	"github.com/sheodox/seating-chart/interactors"
)

type LineController struct {
	Interactor interactors.LineInteractor
}

type NewLineParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	People    int    `json:"people"`
}

type LineCoordinateDTO struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type LineDTO struct {
	Id     string              `json:"id"`
	Coords []LineCoordinateDTO `json:"coords"`
	Closed bool                `json:"closed"`
}

func lineEntityeToDTO(ent entities.Line) LineDTO {
	coords := make([]LineCoordinateDTO, 0)

	for _, coord := range ent.Coords {
		coords = append(coords, LineCoordinateDTO{X: coord.X, Y: coord.Y})
	}

	return LineDTO{
		Id:     ent.Id,
		Closed: ent.Closed,
		Coords: coords,
	}
}

func (l LineController) List() ([]LineDTO, error) {
	lineDto := make([]LineDTO, 0)

	lineEntities, err := l.Interactor.List()

	if err != nil {
		return lineDto, err
	}

	for _, ent := range lineEntities {
		lineDto = append(lineDto, lineEntityeToDTO(ent))
	}

	return lineDto, nil
}

func (l LineController) Add(data map[string]any) (LineDTO, error) {
	lineEntity, err := l.Interactor.Add()

	return lineEntityeToDTO(lineEntity), err
}

func (l LineController) Edit(data map[string]any) (LineDTO, error) {
	// todo safety check
	id := data["id"].(string)
	closed := data["closed"].(bool)
	coords := make([]entities.LineCoordinate, 0)

	if c, ok := data["coords"].([]any); ok {
		for _, point := range c {
			if coord, ok := point.(map[string]any); ok {
				coords = append(coords, entities.LineCoordinate{
					X: coord["x"].(float64),
					Y: coord["y"].(float64),
				})
			}
		}
	}

	lineEntity, err := l.Interactor.Edit(entities.Line{
		Id:     id,
		Coords: coords,
		Closed: closed,
	})

	return lineEntityeToDTO(lineEntity), err
}

func (l LineController) Delete(data map[string]any) (string, error) {
	id, ok := data["id"].(string)

	if !ok {
		return "", ErrBadData
	}

	err := l.Interactor.Delete(id)

	return id, err
}
