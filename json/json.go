package json

import (
    "github.com/toshaf/exhibit/core"
    "encoding/json"
    "fmt"
    "io"
)

func Compare(approved, incoming io.Reader) ([]core.Diff, error) {
    app, err := Load(approved)
    if err != nil {
        return nil, err
    }
    inc, err := Load(incoming)
    if err != nil {
        return nil, err
    }

    return compare(app, inc, nil), nil
}

func compare(app, inc Model, path []string) []core.Diff {
    diffs := []core.Diff{}
    switch a := app.(type) {
    case Object:
        i, is := inc.(Object)
        if !is {
           diffs = append(diffs, core.Diff{
                Expected: a,
                Actual: i,
                Pos: path,
            })
            return diffs
        }

        for ka, va := range a {
            if vi, ok := i[ka]; ok {
                sd := compare(va, vi, append(path, fmt.Sprintf("%v", ka)))
                diffs = append(diffs, sd...)
            } else {
                diffs = append(diffs, core.Diff{
                    Expected: ka,
                    Actual: nil,
                    Pos: path,
                })
            }
        }

        // check the other way
        for ki, _ := range i {
            _, ok := a[ki]
            if !ok {
                diffs = append(diffs, core.Diff{
                    Expected: nil,
                    Actual: ki,
                    Pos: path,
                })
            }
        }
    case Array:
        i, is := inc.(Array)
        if !is {
            diffs = append(diffs, core.Diff{
                Expected: a,
                Actual: i,
                Pos: path,
            })
            return diffs
        }
        if len(a) != len(i) {
            diffs = append(diffs, core.Diff{
                Expected: fmt.Sprintf("Array length %d", len(a)),
                Actual: fmt.Sprintf("Array length %d", len(i)),
                Pos: path,
            })
        }
        for ix, va := range a {
            if ix > len(i) - 1 {
                break
            }

            vi := i[ix]

            diffs = append(diffs, compare(va, vi, append(path, fmt.Sprintf("[%d]", ix)))...)
        }
    default:
        if app != inc {
            diffs = append(diffs, core.Diff{
                Expected: app,
                Actual: inc,
                Pos: path,
            })
        }
    }

    return diffs
}

type Error string

func (e Error) Error() string {
    return string(e)
}

func Errorf(format string, args ...interface{}) Error {
    return Error(fmt.Sprintf(format, args...))
}

// Object, Array or Value
type Model interface{}

// bool, float64, json.Number, string or nil
type Value interface{}

type Object map[Value]Model

type Array []Model

func Load(rdr io.Reader) (Model, error) {
    dec := json.NewDecoder(rdr)

    v, err := decode(dec)
    if err == nil || err == io.EOF {
        return v, nil
    }

    return nil, err
}

func decode(dec *json.Decoder) (Model, error) {
    tok, err := dec.Token()

    if err != nil {
        return nil, err
    }

    switch token := tok.(type) {
    case json.Delim:
        switch rune(token) {
        case '{':
            return decodeObject(dec)
        case '[':
            return decodeArray(dec)
        default:
            return nil, Errorf("Unexpected token %s", token)
        }

    default:
        return token, nil
    }
}

func decodeObject(dec *json.Decoder) (Object, error) {
    obj := make(Object)
    for {
        key, err := dec.Token()
        if err != nil {
            return nil, err
        }
        
        if delim, is := key.(json.Delim); is {
            switch delim {
            case '}':
                return obj, nil
            default:
                return obj, Errorf("Unexpected delim: %s", delim)
            }
        }

        // key is a Value

        value, err := decode(dec)
        obj[key] = value

        if err != nil {
            return obj, err
        }
    }
}

func decodeArray(dec *json.Decoder) (Array, error) {
    arr := make(Array,0)
    for {
        tok, err := dec.Token()
        if err != nil {
            return arr, err
        }

        switch token := tok.(type) {
        case json.Delim:
            switch rune(token) {
            case ']':
                return arr, nil
            case '{':
                obj, err := decodeObject(dec)
                if err != nil {
                    return arr, err
                }
                arr = append(arr, obj)
            case '[':
                return decodeArray(dec)
            default:
                return nil, Errorf("Unexpected token %s", token)
            }
        }
    }
}

