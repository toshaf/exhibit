package exhibit_test

import (
	"testing"
  . "exhibit"
)

func TestSimpleApprovedValue(t *testing.T) {
	Exhibit{t}.Present(TextString("hi"))
}

func TestLabelledValue(t *testing.T) {
	value := []byte("banana")
	Exhibit{t}.PresentLabelled(Text(value), "a")
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
	Exhibit{t}.Present(JSONObj(people))
}
