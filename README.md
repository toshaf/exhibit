# Exhibit

## Evidence based testing

Writing individual assertions on large data structures can be error-prone, brittle and (most importantly) boring.

If you have known-good data that can be deterministically reproduced you could instead capture it as evidence.

For example, you could test your API endpoints
    
    package endpoint_test

    import (
        "testing"
      . "github.com/toshaf/exhibit"
    )

    func TestYourAPIEndpoints(t *testing.T) {
        
        var payload []byte = // get this from somewhere

        Exhibit.A(JSON(payload), t)
    }

When you run `go test` this call to Exhibit will look for a file in your package directory called `endpoint_test.TestYourAPIEndpoints.exhibit-a.json` and compare the contents to your payload (once they've been pretty-printed). If that file doesn't exist or the contents differ, the test fails.

You can also test multiple exhibits per test method, like this

    package some_test

    import (
        "testing"
      . "github.com/toshaf/exhibit"
    )
    
    func TestMultipleExhibits(t *testing.T) {
        Exhibit.A(TextString("This is the content of exhibit A"), t)
        Exhibit.B(TextString("This is the content of exhibit B"), t)
    }

This will produce the following exhibit files
- `some_test.TestMultipleExhibits.exhibit-a.txt`
- `some_test.TestMultipleExhibits.exhibit-b.txt`

You can create exhibits A through Z (you really shouldn't need more than this).

## Workflow

Does this mean I have to create all these files by hand?

No. By running `go test --snapshot` you can create/replace all exhibit files in the packages covered by the test. You do this either to get started or to update your known-good as your project progresses. Be sure to review the changes made by `--snapshot` and only commit them if they're OK.

You can, of course use existing files or create them by hand.

## Types of evidence

There's a very simple `Evidence` interface in case you want to provide your own way of presenting in value to be tested and a custom file extension, but you should see if there's a suitable type already provided.

### Text

Test evidence is stored in .txt exhibit files and does not changes the data it is given.

You can create text evidence in the following ways:
- from a string using `TextString(string) Evidence`
- from a slice of bytes using `Text([]byte) Evidence`
- from a reader using `TextReader(io.Reader) Evidence`

### JSON

JSON evidence is stored in .json exhibit files. The JSON structure is pretty printed to make diffing easier and to make the evidence file more useful as documentation.

You can create JSON evidence in the following ways:
- from a string using `JSONString(string) Evidence`
- from a slice of bytes using `JSON([]byte) Evidence`
- from reader using `JSONReader(io.Reader) Evidence`
- from an arbitrary object using `JSONObj(interface{}) Evidence`

### XML

XML evidence is stored in .xml exhibit files. The XML structure is pretty printed to make diffing easier and to make the evidence file more useful as documentation.

You can create XML evidence in the following ways:
- from a string using `XMLString(string) Evidence`
- from a slice of bytes using `XML([]byte) Evidence`
- from reader using `XMLReader(io.Reader) Evidence`
- from an arbitrary object using `XMLObj(interface{}) Evidence`

## Installation

Simply install using `go get github.com/toshaf/exhibit` or some such.

## Contribute

Comments and pull requests are most welcome :)

