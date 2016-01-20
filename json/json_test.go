package json_test

import (
    "bytes"
    "github.com/toshaf/exhibit/json"
    "testing"
)

func Test_SimilarJSONObjects(t *testing.T) {
    const (
        a = `{"name":"Jeff","age":55}` 
        b = `
            {
                "name": "Jeff",
                "age": 55
            }`
    )

    abuffer := bytes.NewBufferString(a)
    bbuffer := bytes.NewBufferString(b)

    diffs, err := json.Compare(abuffer, bbuffer)
    if err != nil {
        t.Error(err)
        t.FailNow()
    }

    for _, diff := range diffs {
        t.Errorf(diff.String())
    }
}

func Test_DifferentJSONObjects(t *testing.T) {
    const (
        a = `{"name":"Jeff","age":55}` 
        b = `
            {
                "name": "Jeff",
                "age": 56
            }`
    )

    abuffer := bytes.NewBufferString(a)
    bbuffer := bytes.NewBufferString(b)

    diffs, err := json.Compare(abuffer, bbuffer)
    if err != nil {
        t.Error(err)
        t.FailNow()
    }

    if len(diffs) == 0 {
        t.Errorf("Expected a difference, got none")
        t.FailNow()
    }

    diff := diffs[0]
    if ex, is := diff.Expected.(float64); !is || ex != 55 {
        t.Errorf("Got wrong expected value: %v", diff.Expected)
    }
    if ac, is := diff.Actual.(float64); !is || ac != 56 {
        t.Errorf("Got wrong actual value: %v", diff.Actual)
    }
}
