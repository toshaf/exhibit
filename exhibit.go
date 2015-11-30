package exhibit

import (
	"io"
	"io/ioutil"
	"testing"
)

type Exhibit struct {
	T *testing.T
}

type Evidence interface {
	io.Reader
	Extension() string
}

func (ex Exhibit) Present(evidence Evidence) {
	ex.PresentLabelled(evidence, "")
}

func (ex Exhibit) PresentLabelled(evidence Evidence, label string) {
	t := ex.T

	caller, err := getCallerInfo()
	if err != nil {
		t.Errorf("Stack walk failed: %s", err)
		return
	}

	var value string
	if v, e := ioutil.ReadAll(evidence); e == nil {
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
			if *fixup {
				t.Logf("Fixing up :D")
				ioutil.WriteFile(file, []byte(value), 0755)
			} else {
				t.Logf("Expected '%s' but got '%s'", approved, value)
				t.Error()
			}
		}
	}
}
