package entity

import (
	"fmt"
)

type struct Course {
    Title string,
	Credits int,
	Members []*Student
}