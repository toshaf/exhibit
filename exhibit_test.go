package exhibit_test

import (
	. "exhibit"
	"testing"
)

func TestSimpleApprovedValue(t *testing.T) {
	Exhibit.A(TextString("hi"), t)
}

func TestMultipleExhibits(t *testing.T) {
	Exhibit.A(TextString("This is the content of exhibit A"), t)
	Exhibit.B(TextString("This is the content of exhibit B"), t)
}

type Person struct {
	Name string
	Age  int
}

func TestSomeJson(t *testing.T) {
	people := []Person{
		{"Ann", 38},
		{"Bob", 65},
		{"Jeff", 103},
	}
	Exhibit.A(JSONObj(people), t)
}
