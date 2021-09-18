package grades_test

import (
	"testing"

	"github.com/asahnoln/abay-grades"
)

func TestCreateStudent(t *testing.T) {
	s := grades.NewStudent("Arthur", "Aslanyan")

	want := "Arthur Aslanyan"
	got := s.FullName()

	if want != got {
		t.Errorf("want name %q, got %q", want, got)
	}
}

func TestSudentAddGrade(t *testing.T) {
	s := grades.NewStudent("Arthur", "Aslanyan")

	s.AddGrade(10).AddGrade(8).AddGrade(6)

	want := 8
	got := s.AverageGrade()

	if want != got {
		t.Errorf("want grade %d, got %d", want, got)
	}
}
