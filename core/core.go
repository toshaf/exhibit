package core

import (
    "fmt"
    "strings"
)

type Diff struct {
	Expected interface{}
	Actual   interface{}
	Pos []string
}

func (d *Diff) String() string {
    pos := strings.Join(d.Pos, ".")
	expected := format(d.Expected)
	actual := format(d.Actual)
	return fmt.Sprintf("At %s:\n--- %s\n+++ %s", pos, expected, actual)
}

func format(v interface{}) string {
	if s, is := v.(string); is {
		return fmt.Sprintf(`"%s"`, s)
	}

	return fmt.Sprintf("%v", v)
}

