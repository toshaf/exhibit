package rubber

import (
	"testing"
)

func TestSimpleApprovedValue(t *testing.T) {
	RubberStamp{t}.Stamp("hi")
}
