package exhibit

import (
	"flag"
  "fmt"
	"io/ioutil"
  "io"
	"path"
	"runtime"
  "regexp"
	"testing"
  "strings"
)

var (
  fixup *bool
  whitespace *regexp.Regexp
  testname  *regexp.Regexp
  maxdepth int
)

func init() {
	fixup = flag.Bool("fixup", false, "Fixup failing tests by overwriting the approved content")

	flag.Parse()

  whitespace = regexp.MustCompile(`\s+`)
  testname = regexp.MustCompile(`^.*\.Test[^a-z].*`)
  maxdepth = 12
}

type Exhibit struct {
  T *testing.T
}

type Evidence interface {
  io.Reader
  Extension() string
}

type TextEvidence struct {
  io.Reader
}

func (TextEvidence) Extension() string {
  return "txt"
}

func Text(v string) Evidence {
  return TextEvidence{strings.NewReader(v)}
}

func makeEvidenceFilename(evidence Evidence, caller *callerInfo, label string) string {
  label = strings.TrimSpace(label)
  if len(label) > 0 {
    label = "." + string(whitespace.ReplaceAll([]byte(label), []byte{'_'}))
  }

  name := fmt.Sprintf("%s.exhibit%s.%s", caller.function, label, evidence.Extension())
  dir := path.Dir(caller.file)
  return path.Join(dir, name)
}

func (ex Exhibit) Present(evidence Evidence) {
  ex.PresentLabelled(evidence, "")
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
      return &callerInfo {
        file: file,
        function: caller.Name(),
      }, nil
    }
  }
  return nil, fmt.Errorf("Max stack depth (%d) reached, no test method found", maxdepth)
}

func (ex Exhibit) PresentLabelled(evidence Evidence, label string){
  t := ex.T

  caller, err := getCallerInfo()
	if err != nil {
    t.Errorf("Stack walk failed: %s", err)
    return
  }

  var value string
  if v,e := ioutil.ReadAll(evidence); e == nil {
    value = string(v)
  } else {
    t.Errorf("Could not read content: %s", e)
  }

  file := makeEvidenceFilename(evidence, caller, label)

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
