package grades

type Student struct {
	name, lastname string
	grades         []int
}

func NewStudent(name, lastname string) *Student {
	return &Student{
		name:     name,
		lastname: lastname,
	}
}

func (s *Student) FullName() string {
	return s.name + " " + s.lastname
}

func (s *Student) AddGrade(g int) *Student {
	s.grades = append(s.grades, g)
	return s
}

func (s *Student) AverageGrade() int {
	sum := 0
	for _, g := range s.grades {
		sum += g
	}
	return sum / len(s.grades)
}
