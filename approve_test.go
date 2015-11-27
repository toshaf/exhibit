package approve

import (
	"runtime"
	"testing"
)

func approve(t *testing.T, value string) {
	pc, fname, line, ok := runtime.Caller(1)
	if ok {

    caller := runtime.FuncForPC(pc)
		t.Logf(
			"You were called by %s in '%s', line %d (program counter %d)",
      caller.Name(),
			fname,
			line,
			pc,
		)

    t.Error()
	}
}

func TestRuntimeCaller(t *testing.T) {
	approve(t, "hi")
}
