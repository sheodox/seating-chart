package interactors

import (
	"errors"
)

var (
	ErrInvalidPeople    = errors.New("a guest cannot be less than one person")
	ErrInvalidName      = errors.New("a first and last name must be specified")
	ErrInvalidTableName = errors.New("a name must be specified")
	ErrInvalidId        = errors.New("must include a valid guest ID")
	ErrInvalidCapacity  = errors.New("table capacity must be at least 1")
)

func clampPos(pos, bufferSpace float64) float64 {
	// positions are decimals as percentages. don't let them go off the edge
	if pos < bufferSpace {
		return bufferSpace
	} else if pos > 1-bufferSpace {
		return 1 - bufferSpace
	}
	return pos
}
