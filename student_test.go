package student_test

import (
	"slices"
	"testing"
	"time"

	"github.com/janiceb75/student"
)

func TestListWithNoStudentsReturnsEmptyList(t *testing.T) {
	t.Parallel()
	got := student.List()
	want := []student.Student{}
	if slices.Equal(want, got) {
		t.Errorf("want empty list, but got %v", got)
	}
}

func TestAddAddsStudentToDatabase(t *testing.T) {
	t.Parallel()
	birth, err := time.Parse(time.RFC3339, "2018-07-04")
	if err != nil {
		t.Fatal(err)
	}
	peter := student.Student{
		Name:      "Peter Brown",
		BirthDate: birth,
	}
	student.Add(peter)
	want := []student.Student{peter}
	got := student.List()
	if !slices.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
