package exhibit

import (
	"bytes"
	"encoding/json"
    . "github.com/toshaf/exhibit/json"
    "github.com/toshaf/exhibit/core"
	"io"
	"io/ioutil"
)

type jsonEvidence struct {
	buffer bytes.Buffer
}

func (*jsonEvidence) Extension() string {
	return ".json"
}

func (j *jsonEvidence) Check(approved io.Reader) ([]core.Diff, error) {
    return Compare(approved, &j.buffer)
}

func (j *jsonEvidence) GetValue() ([]byte, error) {
    return j.buffer.Bytes(), nil
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
