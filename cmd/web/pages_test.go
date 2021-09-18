package web_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	grades "github.com/asahnoln/abay-grades"
	"github.com/asahnoln/abay-grades/cmd/web"
)

func TestStudentList(t *testing.T) {
	g := grades.NewGroup(
		grades.NewStudent("Art", "Mahov"),
		grades.NewStudent("Gup", "Borev"),
		grades.NewStudent("Wah", "Duhov"),
	)

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	web.Serve(w, r, g)

	want := http.StatusOK
	got := w.Code

	if want != got {
		t.Errorf("want response status %d, got %d", want, got)
	}
}
