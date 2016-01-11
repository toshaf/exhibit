package exhibit_test

import (
	"encoding/xml"
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

func (people People) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	name := xml.Name{Space: "", Local: "People"}

	enc.EncodeToken(xml.StartElement{Name: name})
	for _, person := range people {
		enc.Encode(person)
	}
	enc.EncodeToken(xml.EndElement{Name: name})

	return nil
}

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

func TestRawXml(t *testing.T) {
    xml, _ := xml.MarshalIndent(people, "", "  ")

    Exhibit.A(XMLFormatted(xml), t)
}
