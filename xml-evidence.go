package exhibit

import (
	"bytes"
	"encoding/xml"
	"github.com/toshaf/exhibit/core"
	"github.com/toshaf/exhibit/text"
	"io"
	"io/ioutil"
)

type xmlEvidence struct {
	bytes.Buffer
}

func (*xmlEvidence) Extension() string {
	return ".xml"
}

func (x *xmlEvidence) Check(approved io.Reader) ([]core.Diff, error) {
	return text.Compare(approved, &x.Buffer)
}

func (x *xmlEvidence) GetValue() ([]byte, error) {
	v, err := ioutil.ReadAll(&x.Buffer)
	if err != io.EOF && err != nil {
		return v, err
	}

	return v, nil
}

func XML(v []byte) Evidence {
	buf, err := indentXml(bytes.NewReader(v))
	if err != nil {
		return writeErrorXml(err)
	}

	buf.Write([]byte{'\n'})

	return &xmlEvidence{buf}
}

func XMLFormatted(v []byte) Evidence {
	buf := bytes.NewBuffer(v)
	buf.Write([]byte("\n"))

	return &xmlEvidence{*buf}
}

func XMLString(v string) Evidence {
	return XML([]byte(v))
}

func XMLObj(v interface{}) Evidence {
	b, err := xml.Marshal(v)
	if err != nil {
		return writeErrorXml(err)
	}

	return XML(b)
}

func XMLReader(r io.Reader) Evidence {
	b, e := ioutil.ReadAll(r)
	if e != nil {
		return writeErrorXml(e)
	}

	return XML(b)
}

func writeErrorXml(err error) Evidence {
	return &xmlEvidence{writeError(err)}
}

func indentXml(rdr io.Reader) (bytes.Buffer, error) {
	dec := xml.NewDecoder(rdr)
	buf := bytes.Buffer{}
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "\t")

	for {
		token, err := dec.Token()
		switch err {
		case io.EOF:
			return buf, nil
		default:
			return buf, err
		case nil:
		}

		enc.EncodeToken(token)
		enc.Flush()
	}
}
