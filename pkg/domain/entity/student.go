package entity

import (
	"fmt"
)

type struct Student {
	Id int,
	FirstName string,
	LastName string,
	Courses []*Courses,
	Score int

}