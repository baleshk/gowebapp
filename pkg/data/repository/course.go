package repository

import (
	"fmt"
)

type courseUpdater interface {
	Add (cour *Course) error
	Remove (cour *Course) error
	Update (cour *Course) error
	Get (id int) error
	GetAll () error
}