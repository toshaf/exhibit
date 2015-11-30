package exhibit

import (
	"fmt"
	"regexp"
	"strings"
    "path"
)

var whitespace *regexp.Regexp = regexp.MustCompile(`\s+`)

func makeEvidenceFilename(evidence Evidence, label string) (string, error) {
	caller, err := getCallerInfo()
	if err != nil {
		return "", err
	}

	label = strings.TrimSpace(label)
	if len(label) > 0 {
		label = "-" + string(whitespace.ReplaceAllString(label, "_"))
	}

	name := fmt.Sprintf("%s.exhibit%s%s", caller.function, label, evidence.Extension())

	return path.Join(path.Dir(caller.file), name), nil
}
