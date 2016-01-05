package exhibit

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

type xmlEvidence struct {
	bytes.Buffer
}

//Extension will return the extension of a xml file '.xml'
func (*xmlEvidence) Extension() string {
	return ".xml"
}

//XML takes a byte array and will return a json Evidence object.
func XML(b []byte) Evidence {
	//This is how I think you can format the xml to make it more readable.
	//not sure how to do it from a byte slice..
	//v = the object(s)
	var buff bytes.Buffer
	// encoder := xml.NewEncoder(&buff)
	// encoder.Indent("", "   ")
	// encoder.Encode(v)

	//for now write as a string
	buff.Write(b)
	buff.Write([]byte{'\n'})

	return &xmlEvidence{buff}
}

//XMLObj will marshal the given object to json and return a json Evidence object.
func XMLObj(v interface{}) Evidence {
	var b []byte
	var e error
	if b, e = xml.Marshal(v); e != nil {
		return writeXmlError(e)
	}

	return XML(b)
}

//XMLString takes a string and returns a json Evidence object.
func XMLString(v string) Evidence {
	return XML([]byte(v))
}

//XMLReader takes a reader interface that will be read and return a json Evidence object.
func XMLReader(r io.Reader) Evidence {
	b, e := ioutil.ReadAll(r)
	if e != nil {
		return writeXmlError(e)
	}

	return XML(b)
}

func writeXmlError(e error) Evidence {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "@ %s -> %s", time.Now(), e)

	return &xmlEvidence{buff}
}
