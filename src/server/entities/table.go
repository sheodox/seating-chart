package entities

type Table struct {
	Id       string  `db:"id"`
	Name     string  `db:"name"`
	Capacity int     `db:"capacity"`
	PosX     float64 `db:"pos_x"`
	PosY     float64 `db:"pos_y"`
	Notes    string  `db:"notes"`
}
