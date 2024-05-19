package options_test

import (
	"testing"

	"github.com/janiceb75/students/options"
)

func TestFullName(t *testing.T) {
	s := options.Student{
		FirstName: "Janice",
		LastName:  "Bailey",
	}
	name := s.FullName()
	want := "Janice Bailey"
	if name != want {
		t.Error("You failed. Wanted", want, "Got", name)

	}

}
