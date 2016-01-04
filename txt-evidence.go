package exhibit

import (
	"bytes"
	"io"
)

type textEvidence struct {
	io.Reader
}

//Extension will return txt, the extension for text files.
func (textEvidence) Extension() string {
	return ".txt"
}

//TextString takes a string and will return a text Evidence object.
func TextString(v string) Evidence {
	return Text([]byte(v))
}

//Text will take a text byte array and return a text Evidence object.
func Text(v []byte) Evidence {
	b := bytes.NewBuffer(v)

	return TextReader(b)
}

//TextReader takes a Reader and will return a txt Evidence object.
func TextReader(r io.Reader) Evidence {
	var b bytes.Buffer
	io.Copy(&b, r)
	b.Write([]byte{'\n'})

	return textEvidence{&b}
}
