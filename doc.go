/*
Package exhibit provides a way of using some text based files (txt json at the moment) and
being able to create a known assertion based off that evidence. It hooks into go test and
will fail if the output does not match the original assertion which you must deem as correct.

You also have a way of overwritting the assertion if it does infact need to be updated.

Examples of how to use:
  package endpoint_test

  import (
      "testing"
    . "github.com/toshaf/exhibit"
  )

  func TestYourAPIEndpoints(t *testing.T) {

      var payload []byte = // get this from somewhere

      Exhibit.A(JSON(payload), t)
  }

or

  package some_test

  import (
      "testing"
    . "github.com/toshaf/exhibit"
  )

  func TestMultipleExhibits(t *testing.T) {
      Exhibit.A(TextString("This is the content of exhibit A"), t)
      Exhibit.B(TextString("This is the content of exhibit B"), t)
  }
*/
package exhibit
