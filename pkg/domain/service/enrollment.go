package service

import (
	"fmt"
)

func Enroll (stu *Student, cour *Course) error{
	if cour.Credits>stu.Score {
		return errors.New("math: divided by zero")
	}
	stu.Courses=append(stu.Courses, cour)
	cour.Members=append(cour.Members, cour)
}