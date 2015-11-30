package exhibit

import (
	"fmt"
	"regexp"
	"runtime"
)

const maxdepth int = 12

var testname *regexp.Regexp = regexp.MustCompile(`^.*\.Test[^a-z].*`)

type callerInfo struct {
	file, function string
}

func getCallerInfo() (*callerInfo, error) {

	for i := 2; i < maxdepth; i++ {

		pc, file, _, ok := runtime.Caller(i)
		if !ok {
			return nil, fmt.Errorf("Could not retrieve caller %d", i)
		}

		caller := runtime.FuncForPC(pc)

		if testname.MatchString(caller.Name()) {
			return &callerInfo{
				file:     file,
				function: caller.Name(),
			}, nil
		}
	}

	return nil, fmt.Errorf("Max stack depth (%d) reached, no test method found", maxdepth)
}
