package exhibit

import (
	"flag"
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

var (
	fixup      *bool
	whitespace *regexp.Regexp
	testname   *regexp.Regexp
	prefix     *regexp.Regexp
	maxdepth   int
)

func init() {
	fixup = flag.Bool("fixup", false, "Fixup failing tests by overwriting the approved content")

	flag.Parse()

	whitespace = regexp.MustCompile(`\s+`)
	testname = regexp.MustCompile(`^.*\.Test[^a-z].*`)
	prefix = regexp.MustCompile(`(^.*\.)?`)
	maxdepth = 12
}

func makeEvidenceFilename(evidence Evidence, caller *callerInfo, label string) string {
	label = strings.TrimSpace(label)
	if len(label) > 0 {
		label = "-" + string(whitespace.ReplaceAllString(label, "_"))
	}

	name := fmt.Sprintf("%s.exhibit%s.%s", prefix.ReplaceAllString(caller.function, ""), label, evidence.Extension())
	return fmt.Sprintf("%s.%s", caller.file, name)
}

type callerInfo struct {
	file, function string
}

func getCallerInfo() (*callerInfo, error) {
	for i := 2; i < maxdepth; i++ {
		// program counter, filename, line, ok
		pc, file, _, ok := runtime.Caller(i)
		if !ok {
			// return nil, fmt.Errorf("Could not retrieve caller %d", i)
			continue
		}

		caller := runtime.FuncForPC(pc)

		if testname.Match([]byte(caller.Name())) {
			return &callerInfo{
				file:     file,
				function: caller.Name(),
			}, nil
		}
	}
	return nil, fmt.Errorf("Max stack depth (%d) reached, no test method found", maxdepth)
}
