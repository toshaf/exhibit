package exhibit

import "flag"

var snapshot *bool

func init() {
	snapshot = flag.Bool("snapshot", false, "Snapshot all evidence, skips dependent tests")

	flag.Parse()
}
