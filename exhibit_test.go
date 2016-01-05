package exhibit_test

import (
	. "github.com/toshaf/exhibit"
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

type People []Person

var people = People{
	{"Ann", 38},
	{"Bob", 65},
	{"Jeff", 103},
}

func TestSomeJson(t *testing.T) {
	Exhibit.A(JSONObj(people), t)
}

func TestSomeXml(t *testing.T) {
	Exhibit.A(XMLObj(people), t)
}
