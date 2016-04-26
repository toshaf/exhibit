package exhibit

import (
	"bytes"
	"github.com/toshaf/exhibit/core"
	"github.com/toshaf/exhibit/text"
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
	return text.Compare(approved, t.Reader)
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

	return textEvidence{&b}
}
