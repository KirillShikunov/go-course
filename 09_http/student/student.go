package student

var students = []*Student{
	NewStudent(1, "John Doe", 1, map[string]float64{"math": 90, "science": 80}),
	NewStudent(2, "Alex Turner", 1, map[string]float64{"math": 85, "science": 95}),
	NewStudent(3, "Michael Smith", 2, map[string]float64{"math": 70, "science": 75}),
	NewStudent(4, "Maria Garcia", 2, map[string]float64{"math": 60, "science": 65}),
}

type Student struct {
	Id        int
	Name      string
	ClassId   int
	AvgScores map[string]float64
}

func NewStudent(id int, name string, classId int, avgScore map[string]float64) *Student {
	return &Student{id, name, classId, avgScore}
}

func Find(id int) *Student {
	for _, s := range students {
		if s.Id == id {
			return s
		}
	}

	return nil
}

func FindStudentsByClass(classId int) []*Student {
	var foundStudents []*Student
	for _, s := range students {
		if s.ClassId == classId {
			foundStudents = append(foundStudents, s)
		}
	}

	return foundStudents
}
