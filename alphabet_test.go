package exhibit_test

import (
  "testing"
  . "exhibit"
)

func Test_A_method(t *testing.T){
  Exhibit{t}.A(TextString("This is the content of exhibit A"))
}

