package exhibit

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

type jsonEvidence struct {
	bytes.Buffer
}

func (*jsonEvidence) Extension() string {
	return ".json"
}

func JSON(v []byte) Evidence {
	var buff bytes.Buffer
	if e := json.Indent(&buff, v, "", "\t"); e != nil {
		return writeErrorJson(e)
	}

	buff.Write([]byte{'\n'})

	return &jsonEvidence{buff}
}

func JSONString(v string) Evidence {
	return JSON([]byte(v))
}

func JSONObj(v interface{}) Evidence {
	if b, e := json.Marshal(v); e != nil {
		return writeErrorJson(e)
	} else {
		return JSON(b)
	}
}

func JSONReader(r io.Reader) Evidence {
	b, e := ioutil.ReadAll(r)
	if e != nil {
		return writeErrorJson(e)
	}

	return JSON(b)
}

func writeErrorJson(e error) *jsonEvidence {
	return &jsonEvidence{writeError(e)}
}
