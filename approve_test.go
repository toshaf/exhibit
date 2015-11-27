package approve

import (
	"runtime"
	"testing"
)

func approve(t *testing.T, value string) {
	pc, fname, line, ok := runtime.Caller(1)
	if ok {
		t.Logf(
			"You were called by '%s', line %d (program counter %d)",
			fname,
			line,
			pc,
		)
	}
  t.Error()
}

func TestRuntimeCaller(t *testing.T) {
	approve(t, "hi")
}
