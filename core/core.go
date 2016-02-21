package core

import (
	"fmt"
	"strings"
)

type Diff struct {
	Expected interface{}
	Actual   interface{}
	Pos      []string
}

func (d Diff) Equals(e Diff) bool {
	if d.Expected != e.Expected {
		return false
	}
	if d.Actual != e.Actual {
		return false
	}
	for i, pos := range d.Pos {
		if e.Pos[i] != pos {
			return false
		}
	}
	return true
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
