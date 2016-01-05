package exhibit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"
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
		return writeJsonError(e)
	}

	buff.Write([]byte{'\n'})

	return &jsonEvidence{buff}
}

func JSONString(v string) Evidence {
	return JSON([]byte(v))
}

func JSONObj(v interface{}) Evidence {
	var b []byte
	var e error
	if b, e = json.Marshal(v); e != nil {
		return writeJsonError(e)
	}

	return JSON(b)
}

func JSONReader(r io.Reader) Evidence {
	b, e := ioutil.ReadAll(r)
	if e != nil {
		return writeJsonError(e)
	}

	return JSON(b)
}

func writeJsonError(e error) Evidence {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "@ %s -> %s", time.Now(), e)

	return &jsonEvidence{buff}
}
