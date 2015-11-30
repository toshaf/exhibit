package exhibit

import (
	"io"
	"strings"
)

type textEvidence struct {
	io.Reader
}

func (textEvidence) Extension() string {
	return "txt"
}

func TextString(v string) Evidence {
	return textEvidence{strings.NewReader(v)}
}

func Text(v []byte) Evidence {
	return TextString(string(v))
}
