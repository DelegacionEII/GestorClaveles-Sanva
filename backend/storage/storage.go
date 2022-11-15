package storage

import (
	"github.com/Error404UsernameNotFound/sanva/backend"
)

type storage interface {
	// Adds a new student to the DB, identified by name and UO.
	AddStudent(name string, uo string) error

	// Deletes the student identified by UO.
	DeleteStudent(uo string) error

	// Add a number of flowers to the student identified by UO.
	// Flowers with no color parameter will be given a default one.
	AddFlowers(uo string, number int64, color string) error

	// Returns all flowers associated to a given student,
	// identified by UO. Flowers are returned using a map,
	// with the color as a key and the number as the value.
	RetrieveFlowers(uo string) (backend.Flowers, error)

	// Deletes all current data.
	Clear() error
}
