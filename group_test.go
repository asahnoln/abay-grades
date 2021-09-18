package grades_test

import (
	"testing"

	"github.com/asahnoln/abay-grades"
)

func TestStudentGroup(t *testing.T) {
	sg := grades.NewGroup(
		grades.NewStudent("Abay", "Amantayev"),
		grades.NewStudent("Arthur", "Aslanyan"))
	sg.Add(grades.NewStudent("Kolya", "Mitkin"))

	want := 3
	got := sg.Len()

	if want != got {
		t.Fatalf("want group length %d, got %d", want, got)
	}

	wantName := "Kolya Mitkin"
	gotName := sg.Student(2).FullName()
	if wantName != gotName {
		t.Errorf("want last student name %q, got %q", wantName, gotName)
	}
}
