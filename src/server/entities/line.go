package entities

type LineCoordinate struct {
	X float64 `db:"x"`
	Y float64 `db:"y"`
}

type Line struct {
	Id     string           `db:"id"`
	Coords []LineCoordinate `db:"coords"`
	Closed bool             `db:"closed"`
}
