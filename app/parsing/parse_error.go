package parsing

import (
	"errors"
	"fmt"
)

// errTimeout - to long answer from the site
type errTimeout struct {
	link string
}

func (e errTimeout) Error() string {
	return fmt.Sprintf("Timeout error. Could not reach the site: %s", e.link)
}

// errStatusCode - bad answer from the site
type errStatusCode struct {
	link string
	code int
}

func (e errStatusCode) Error() string {
	return fmt.Sprintf("Bad answer from %s. Status code: %d", e.link, e.code)
}

// errParseBody - error from parsing the body of the answer
type errParseBody struct {
	link string
}

func (e errParseBody) Error() string {
	return fmt.Sprintf("Could not parse the body of the answer: %s", e.link)
}

// errUrlParse - error from parsing incoming url
var errUrlParse = errors.New("Could not parse incoming url")

// errParseJson - incorrect incomming json request
type errParseJson struct {
}

func (e errParseJson) Error() string {
	return fmt.Sprintf("Incorrect incomming JSON")
}
