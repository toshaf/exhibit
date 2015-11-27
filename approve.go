package approve

import (
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func approve(t *testing.T, value string) {
  // program counter, filename, line, ok
	pc, fname, _, ok := runtime.Caller(1)
	if ok {

		caller := runtime.FuncForPC(pc)

		dir := path.Dir(fname)
		file := path.Join(dir, caller.Name() + ".approved")

		if approved, err := ioutil.ReadFile(file); err != nil {
			t.Logf("Could not read approved value from '%s'", file)
			ioutil.WriteFile(file, []byte(value), 0755)
			t.Error()
			return
		} else {
			if value != string(approved) {
				t.Logf("Expected '%s' but got '%s'", approved, value)
				t.Error()
			}
		}
	}
}
