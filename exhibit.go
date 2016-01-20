package exhibit

import (
	"github.com/toshaf/exhibit/core"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type E struct{}

var Exhibit E

type Evidence interface {
	Extension() string
	Check(approved io.Reader) ([]core.Diff, error)
	GetValue() ([]byte, error)
}

func (E) present(evidence Evidence, label string, t *testing.T) {
	file, err := makeEvidenceFilename(evidence, label)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if *snapshot {
		value, err := evidence.GetValue()
		if err != io.EOF && err != nil {
			t.Error(err)
			t.FailNow()
		}
		ioutil.WriteFile(file, []byte(value), 0644)
		t.Logf("Writing Exhibit %s snapshot to %s", label, file)
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
