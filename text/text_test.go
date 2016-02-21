package text_test

import (
	"bytes"
	"github.com/toshaf/exhibit/core"
	"github.com/toshaf/exhibit/text"
	"testing"
)

func TestSimilarStrings(t *testing.T) {
	s := "one\ntwo\nthree\nfour\n"
	a, b := bytes.NewBufferString(s), bytes.NewBufferString(s)

	diffs, err := text.Compare(a, b)

	if err != nil {
		t.Error(err)
	}
	for _, diff := range diffs {
		t.Errorf(diff.String())
	}
}

func TestSimilarStringsDifferentLineEndings(t *testing.T) {
	s1 := "one\ntwo\nthree\nfour\n"
	s2 := "one\r\ntwo\r\nthree\r\nfour\r\n"
	a, b := bytes.NewBufferString(s1), bytes.NewBufferString(s2)

	diffs, err := text.Compare(a, b)

	if err != nil {
		t.Error(err)
	}
	for _, diff := range diffs {
		t.Errorf(diff.String())
	}
}

func TestDifferentStrings(t *testing.T) {
	s1 := "one\ntwo\nthree\nfour\n"
	s2 := "one\nTwo\nthree\nfore\n"
	a, b := bytes.NewBufferString(s1), bytes.NewBufferString(s2)

	diffs, err := text.Compare(a, b)

	if err != nil {
		t.Error(err)
	}
	if len(diffs) != 2 {
		t.Errorf("Expected 2 diffs, got %d", len(diffs))
	}

	if len(diffs) > 0 {
		ed := core.Diff{Expected: "two", Actual: "Two", Pos: []string{"2"}}
		ad := diffs[0]

		if !ad.Equals(ed) {
			t.Errorf("Expected %s but got %s", ed, ad)
		}
	}

	if len(diffs) > 1 {
		ed := core.Diff{Expected: "four", Actual: "fore", Pos: []string{"4"}}
		ad := diffs[1]

		if !ad.Equals(ed) {
			t.Errorf("Expected %s but got %s", ed, ad)
		}
	}
}

func TestLongerThanExpectedString(t *testing.T) {
	s1 := "one\ntwo\nthree\n"
	s2 := "one\ntwo\nthree\nfour\n"
	a, b := bytes.NewBufferString(s1), bytes.NewBufferString(s2)

	diffs, err := text.Compare(a, b)

	if err != nil {
		t.Error(err)
	}

	if len(diffs) != 1 {
		t.Errorf("Expected 1 diffs, got %d", len(diffs))
	}

	if len(diffs) > 0 {
		ed := core.Diff{Expected: nil, Actual: "four", Pos: []string{"4"}}
		ad := diffs[0]

		if !ad.Equals(ed) {
			t.Errorf("Expected %s but got %s", ed, ad)
		}
	}
}

func TestShorterThanExpectedString(t *testing.T) {
	s1 := "one\ntwo\nthree\nfour\n"
	s2 := "one\ntwo\nthree\n"
	a, b := bytes.NewBufferString(s1), bytes.NewBufferString(s2)

	diffs, err := text.Compare(a, b)

	if err != nil {
		t.Error(err)
	}

	if len(diffs) != 1 {
		t.Errorf("Expected 1 diffs, got %d", len(diffs))
		for _, d := range diffs {
			t.Errorf(d.String())
		}
	}

	if len(diffs) > 0 {
		ed := core.Diff{Expected: "four", Actual: nil, Pos: []string{"4"}}
		ad := diffs[0]

		if !ad.Equals(ed) {
			t.Errorf("Expected %s but got %s", ed, ad)
		}
	}
}
