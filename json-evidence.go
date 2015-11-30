package exhibit

import (
  "bytes"
  "fmt"
  "time"
  "encoding/json"
)

type jsonEvidence struct {
  bytes.Buffer
}

func (*jsonEvidence) Extension() string {
  return "json"
}

func JSON(v []byte) Evidence {
  var buff bytes.Buffer
  if e := json.Indent(&buff, v, "", "\t"); e != nil {
    return writeError(e)
  }

  return &jsonEvidence{buff}
}

func JSONString(v string) Evidence {
  return JSON([]byte(v))
}

func JSONObj(v interface{}) Evidence {
  if b, e := json.Marshal(v); e != nil {
    return writeError(e)
  } else {
    return JSON(b)
  }
}

func writeError(e error) Evidence {
  var buff bytes.Buffer
  fmt.Fprintf(&buff, "@ %s -> %s", time.Now(), e)

  return &jsonEvidence{buff}
}
