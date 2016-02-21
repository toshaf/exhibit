package text

import (
	"bufio"
	"fmt"
	"github.com/toshaf/exhibit/core"
	"io"
)

func Compare(a, b io.Reader) ([]core.Diff, error) {
	sa, sb := bufio.NewScanner(a), bufio.NewScanner(b)

	line := 0
	diffs := []core.Diff{}
	for sa.Scan() {
		line++
		if !sb.Scan() {
			diffs = append(diffs, core.Diff{
				Expected: sa.Text(),
				Actual:   nil,
				Pos:      []string{fmt.Sprintf("%d", line)},
			})
			if err := sb.Err(); err != nil {
				return diffs, err
			}
		} else if sa.Text() != sb.Text() {
			diffs = append(diffs, core.Diff{
				Expected: sa.Text(),
				Actual:   sb.Text(),
				Pos:      []string{fmt.Sprintf("%d", line)},
			})
		}
	}

	if err := sa.Err(); err != nil {
		return diffs, err
	}

	for sb.Scan() {
		line++
		diffs = append(diffs, core.Diff{
			Expected: nil,
			Actual:   sb.Text(),
			Pos:      []string{fmt.Sprintf("%d", line)},
		})
	}

	if err := sb.Err(); err != nil {
		return diffs, err
	}

	return diffs, nil
}
