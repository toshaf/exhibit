package exhibit

import "flag"

var fixup *bool

func init() {
	fixup = flag.Bool("fixup", false, "Fixup failing tests by overwriting the approved content")

	flag.Parse()
}
