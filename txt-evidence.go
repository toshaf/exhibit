package exhibit

import (
	"bytes"
	"io"
)

type textEvidence struct {
	io.Reader
}

func (textEvidence) Extension() string {
	return ".txt"
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
