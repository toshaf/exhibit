package exhibit

import (
	"io"
	"io/ioutil"
	"testing"
)

type E struct{}

var Exhibit E

type Evidence interface {
	io.Reader
	Extension() string
}

func (E) present(evidence Evidence, label string, t *testing.T) {
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

	if approved, err := ioutil.ReadFile(file); err != nil {
		t.Errorf("Could not read evidence from file '%s'", file)
	} else if value != string(approved) {
		t.Errorf("Expected '%s' but got '%s'", approved, value)
	}

	if *snapshot {
		ioutil.WriteFile(file, []byte(value), 0644)
		t.Logf("Writing Exhibit %s snapshot to %s", label, file)
		return
	}
}
