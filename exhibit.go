package exhibit

import (
	"github.com/toshaf/exhibit/core"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type E struct{}

var Exhibit E

type Evidence interface {
	Extension() string
	Check(approved io.Reader) ([]core.Diff, error)
	GetValue() ([]byte, error)
}

func (e E) present(evidence Evidence, label string, t *testing.T) {
	file, err := makeEvidenceFilename(evidence, label)
	if err != nil {
		t.Error(err)
		return
	}

	e.Named(file, evidence, t)
}

func (E) Named(file string, evidence Evidence, t *testing.T){
	if *snapshot {
		value, err := evidence.GetValue()
		if err != io.EOF && err != nil {
			t.Error(err)
			t.FailNow()
		}
		err = os.MkdirAll(filepath.Dir(file), os.ModeDir|os.ModePerm)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Writing Exhibit snapshot to %s", file)
		err = ioutil.WriteFile(file, []byte(value), 0644)
		if err != nil {
			t.Error(err)
		}
		return
	}

	approved, err := os.Open(file)
	if err != nil {
		t.Errorf("Could not read evidence from file '%s'", file)
	}

	diffs, err := evidence.Check(approved)
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, diff := range diffs {
		t.Errorf(diff.String())
	}
}
