package exhibit_test

import (
	"testing"
  . "exhibit"
)

func TestSimpleApprovedValue(t *testing.T) {
	Exhibit{t}.A(TextString("hi"))
}

func Test_alphbet_methods(t *testing.T){
  exhibit := Exhibit{t}

  exhibit.A(TextString("This is the content of exhibit A"))
  exhibit.B(TextString("This is the content of exhibit B"))
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
	Exhibit{t}.A(JSONObj(people))
}
