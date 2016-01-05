package exhibit_test

import (
	"testing"

	. "github.com/toshaf/exhibit"
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

type Monkey struct {
	Name           string
	Bananas        int
	AfraidOfTarzan bool
}

func TestSomeXml(t *testing.T) {
	troop := []Monkey{
		{"Bubbles", 5, true},
		{"King Kong", 1000, false},
	}

	Exhibit.A(XMLObj(troop), t)
}
