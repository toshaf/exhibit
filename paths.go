package exhibit

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

var whitespace = regexp.MustCompile(`\s+`)

func makeEvidenceFilename(evidence Evidence, label string) (string, error) {
	caller, err := getCallerInfo()
	if err != nil {
		return "", err
	}

	label = strings.TrimSpace(label)
	if len(label) > 0 {
		label = "-" + string(whitespace.ReplaceAllString(label, "_"))
	}

	_, function := path.Split(caller.function)

	name := fmt.Sprintf("%s.exhibit%s%s", function, label, evidence.Extension())

	return path.Join(path.Dir(caller.file), name), nil
}
