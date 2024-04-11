package entites

import "errors"

var ErrTeacherNotFound = errors.New("teacher not found")

var teachers = []*Teacher{
	NewTeacher(1, "Mr. Smith"),
	NewTeacher(2, "Mr. Doe"),
}

type Teacher struct {
	Id   int
	Name string
}

func NewTeacher(id int, name string) *Teacher {
	return &Teacher{id, name}
}

func FindTeacher(id int) (*Teacher, error) {
	for _, t := range teachers {
		if t.Id == id {
			return t, nil
		}
	}

	return nil, ErrTeacherNotFound
}
