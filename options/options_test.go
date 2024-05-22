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

func TestStudentAge(t *testing.T) {
	a := options.Student{
		Age: 48,
	}

	studentAge := a.Age
	want := 46
	if studentAge != want {
		t.Error("Test failed. I wanted", want, "Got", studentAge)
	}

}
