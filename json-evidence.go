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

//Extension will return the extension of a josn file '.json'
func (*jsonEvidence) Extension() string {
	return ".json"
}

//JSON takes a byte array and will return a json Evidence object.
func JSON(v []byte) Evidence {
	var buff bytes.Buffer
	if e := json.Indent(&buff, v, "", "\t"); e != nil {
		return writeError(e)
	}

	buff.Write([]byte{'\n'})

	return &jsonEvidence{buff}
}

//JSONString takes a string and returns a json Evidence object.
func JSONString(v string) Evidence {
	return JSON([]byte(v))
}

//JSONObj will marshal the given object to json and return a json Evidence object.
func JSONObj(v interface{}) Evidence {
	if b, e := json.Marshal(v); e != nil {
		return writeError(e)
	} else {
		return JSON(b)
	}
}

//JSONReader takes a reader interface that will be read and return a json Evidence object.
func JSONReader(r io.Reader) Evidence {
	b, e := ioutil.ReadAll(r)
	if e != nil {
		return writeError(e)
	}

	return JSON(b)
}

func writeError(e error) Evidence {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "@ %s -> %s", time.Now(), e)

	return &jsonEvidence{buff}
}
