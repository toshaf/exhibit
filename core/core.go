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
	return fmt.Sprintf("At %s:\n--- %s\n+++%s", pos, d.Expected, d.Actual)
}

