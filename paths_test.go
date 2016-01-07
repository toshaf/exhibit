package exhibit

import (
	"testing"
)

func Test_functionSantise(t *testing.T) {
    s := functionSanitise("pkgname.TestSomeStuff.func1.1") 
	if s != "pkgname.TestSomeStuff" {
		t.Error("Didn't replace guff: " + s)
	}

    s = functionSanitise("pkgname.TestSomeStuff")

	if s != "pkgname.TestSomeStuff" {
		t.Error("Didn't leave original alone: " + s)
	}

    s = functionSanitise("TestSomeStuff")

	if s != "TestSomeStuff" {
		t.Error("Didn't handle unqualified name: " + s)
	}
}
