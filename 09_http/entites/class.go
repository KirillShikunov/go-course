package entites

var classes = []*Class{
	NewClass(1, "10-B", NewTeacher(1, "Mr. Smith")),
	NewClass(2, "10-A", NewTeacher(2, "Mr. Doe")),
}

type Class struct {
	Id       int
	Name     string
	Teacher  *Teacher
	Students []*Student
}

func NewClass(id int, name string, teacher *Teacher) *Class {
	return &Class{id, name, teacher, []*Student{}}
}

func FindClass(id int) *Class {
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
