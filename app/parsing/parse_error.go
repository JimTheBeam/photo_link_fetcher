package parsing

import (
	"errors"
	"fmt"
)

// errGetHTML - to long answer from the site
type errGetHTML struct {
	link string
}

func (e errGetHTML) Error() string {
	return fmt.Sprintf("Could not reach the site: %s", e.link)
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

// ErrUrlParse - error from parsing incoming url
var ErrUrlParse = errors.New("Could not parse an incoming url")

// ErrParseJson - incorrect incomming json request
var ErrParseJson = errors.New("Incorrect incomming json")

// ErrEmptyJson - incorrect incomming json request
var ErrEmptyJson = errors.New("Empty array of an incoming url")
