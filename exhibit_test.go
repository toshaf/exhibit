package exhibit

import (
	"testing"
)

func Test_TestNameRegex(t *testing.T){
  if testname.Match([]byte("exhibit.TestSimpleApprovedValue")) == false {
    t.Errorf("Should have matched exhibit.TestSimpleApprovedValue")
  }
}

func TestSimpleApprovedValue(t *testing.T) {
	Exhibit{t}.Present(Text("hi"))
}

func TestLabelledValue(t *testing.T) {
  Exhibit{t}.PresentLabelled(Text("banana"), "a")
}
