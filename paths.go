package exhibit

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	whitespace *regexp.Regexp = regexp.MustCompile(`\s+`)
	prefix     *regexp.Regexp = regexp.MustCompile(`(^.*\.)?`)
)

func makeEvidenceFilename(evidence Evidence, label string) (string, error) {
	caller, err := getCallerInfo()
	if err != nil {
		return "", err
	}

	label = strings.TrimSpace(label)
	if len(label) > 0 {
		label = "-" + string(whitespace.ReplaceAllString(label, "_"))
	}

	function := prefix.ReplaceAllString(caller.function, "")

	name := fmt.Sprintf("%s.exhibit%s%s", function, label, evidence.Extension())

	return fmt.Sprintf("%s.%s", caller.file, name), nil
}
