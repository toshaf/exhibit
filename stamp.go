package rubber

import (
	"flag"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

var fixup *bool

func init() {
	fixup = flag.Bool("fixup", false, "Fixup failing tests by overwriting the approved content")

	flag.Parse()
}

func stamp(t *testing.T, value string) {
	// program counter, filename, line, ok
	pc, fname, _, ok := runtime.Caller(1)
	if ok {

		caller := runtime.FuncForPC(pc)

		dir := path.Dir(fname)
		file := path.Join(dir, caller.Name()+".stamped")

		if approved, err := ioutil.ReadFile(file); err != nil {
			t.Logf("Could not read approved value from '%s'", file)
			if *fixup {
				t.Logf("Fixing up :D")
				ioutil.WriteFile(file, []byte(value), 0755)
			} else {
				t.Error()
			}
			return
		} else {
			if value != string(approved) {
				t.Logf("Expected '%s' but got '%s'", approved, value)
				t.Error()
			}
		}
	}
}
