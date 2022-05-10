package interactors

import (
	"errors"
)

var (
	ErrInvalidPeople    = errors.New("a guest cannot be less than one person")
	ErrInvalidName      = errors.New("a first and last name must be specified")
	ErrInvalidTableName = errors.New("a name must be specified")
	ErrInvalidId        = errors.New("must include a valid guest ID")
)
