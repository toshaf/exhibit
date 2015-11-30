package exhibit

import (
	"io"
	"io/ioutil"
	"testing"
)

type Exhibit struct {
	Testing *testing.T
}

type Evidence interface {
	io.Reader
	Extension() string
}

func (ex Exhibit) Present(evidence Evidence) {
	ex.PresentLabelled(evidence, "")
}

func (ex Exhibit) PresentLabelled(evidence Evidence, label string) {
	t := ex.Testing

	var value string
	if v, e := ioutil.ReadAll(evidence); e == nil {
		value = string(v)
	} else {
		t.Errorf("Could not read content: %s", e)
	}

	file, err := makeEvidenceFilename(evidence, label)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if *fixup {
		t.Logf(file)
		ioutil.WriteFile(file, []byte(value), 0755)
		return
	}

	if approved, err := ioutil.ReadFile(file); err != nil {
		t.Errorf("Could not read evidence from file '%s'", file)
	} else if value != string(approved) {
		t.Logf("Expected '%s' but got '%s'", approved, value)
		t.Error()
	}
}
