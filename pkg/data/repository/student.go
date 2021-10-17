package repository

import (
	"fmt"
)

type studentUpdater interface {
	Add (cour *Student) error
	Remove (cour *Student) error
	Update (cour *Student) error
	Get (id int) error
	GetAll () error
}