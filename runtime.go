package exhibit

import (
	"fmt"
	"regexp"
	"runtime"
)

var testname = regexp.MustCompile(`^.*\.Test[^a-z].*`)

type callerInfo struct {
	file, function string
}

func getCallerInfo() (*callerInfo, error) {

	for i := 2; ; i++ {

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
}
