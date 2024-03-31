package class

import (
	"09_http/student"
	"09_http/teacher"
)

var classes = []*Class{
	NewClass(1, "10-B", teacher.NewTeacher(1, "Mr. Smith")),
	NewClass(2, "10-A", teacher.NewTeacher(2, "Mr. Doe")),
}

type Class struct {
	Id       int
	Name     string
	Teacher  *teacher.Teacher
	Students []*student.Student
}

func NewClass(id int, name string, teacher *teacher.Teacher) *Class {
	return &Class{id, name, teacher, []*student.Student{}}
}

func Find(id int) *Class {
	for _, c := range classes {
		if c.Id == id {
			return c
		}
	}

	return nil
}

func FindClassesByTeacher(teacherId int) []*Class {
	var foundClasses []*Class
	for _, c := range classes {
		classTeacher := c.Teacher
		if classTeacher.Id == teacherId {
			foundClasses = append(foundClasses, c)
		}
	}

	return foundClasses
}
