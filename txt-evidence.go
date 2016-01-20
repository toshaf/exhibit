package exhibit

import (
	"bytes"
	"github.com/toshaf/exhibit/core"
	"io"
	"io/ioutil"
)

type textEvidence struct {
	io.Reader
}

func (textEvidence) Extension() string {
	return ".txt"
}

func (t textEvidence) Check(approved io.Reader) ([]core.Diff, error) {
	diffs := []core.Diff{}

	v1, err := t.GetValue()
	if err != nil {
		return diffs, err
	}

	v2, err := ioutil.ReadAll(approved)
	if err != nil {
		return diffs, err
	}

	if string(v1) != string(v2) {
		diffs = append(diffs, core.Diff{
			Expected: string(v1),
			Actual:   string(v2),
			Pos:      []string{"SIMPLE TXT"},
		})
	}

	return diffs, nil
}

func (t textEvidence) GetValue() ([]byte, error) {
	v, err := ioutil.ReadAll(t.Reader)
	if err != io.EOF && err != nil {
		return nil, err
	}

	return v, nil
}

func TextString(v string) Evidence {
	return Text([]byte(v))
}

func Text(v []byte) Evidence {
	b := bytes.NewBuffer(v)

	return TextReader(b)
}

func TextReader(r io.Reader) Evidence {
	var b bytes.Buffer
	io.Copy(&b, r)
	b.Write([]byte{'\n'})

	return textEvidence{&b}
}
