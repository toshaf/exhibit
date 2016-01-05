package exhibit

import (
	"bytes"
	"fmt"
	"time"
)

func writeError(e error) bytes.Buffer {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "@ %s -> %s", time.Now(), e)

	return buff
}
