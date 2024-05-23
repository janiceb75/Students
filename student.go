package student

import "time"

type Student struct {
	Name      string
	BirthDate time.Time
}

func List() []Student {
	return nil
}

func Add(s Student) {}
