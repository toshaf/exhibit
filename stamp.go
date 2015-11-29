package rubber

import (
	"flag"
	"io/ioutil"
  "io"
	"path"
	"runtime"
	"testing"
  "
)

var fixup *bool

func init() {
	fixup = flag.Bool("fixup", false, "Fixup failing tests by overwriting the approved content")

	flag.Parse()
}

type RubberStamp struct {
  T *testing.T
}

type Content interface {
  io.Reader
}

func String(v string) Content {

}

func (s RubberStamp) Stamp(c Content){
  t := s.T
	// program counter, filename, line, ok
	pc, fname, _, ok := runtime.Caller(1)
	if ok {
    var value string
    if v,e := ioutil.ReadAll(c); e == nil {
      value = string(v)
    } else {
      t.Errorf("Could not read content: %s", e)
    }

		caller := runtime.FuncForPC(pc)

		dir := path.Dir(fname)
		file := path.Join(dir, caller.Name()+".stamped")

		if approved, err := ioutil.ReadFile(file); err != nil {
			t.Logf("Could not read approved value from '%s': %s", file, err)
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
