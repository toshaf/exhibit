package exhibit

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

var whitespace *regexp.Regexp = regexp.MustCompile(`\s+`)

func functionSanitise(name string) string {
	parts := strings.Split(name, ".")
    if len(parts) > 1 {
        return fmt.Sprintf("%s.%s", parts[0], parts[1])
    }

    return parts[0]
}

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

	function = functionSanitise(function)

	name := fmt.Sprintf("%s.exhibit%s%s", function, label, evidence.Extension())

	return path.Join(path.Dir(caller.file), name), nil
}
