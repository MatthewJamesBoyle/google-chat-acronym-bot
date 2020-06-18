package persist

import (
	"errors"
)

type Persister interface {
	LookUpper
	Adder
}

type LookUpper interface {
	LookupPhrase(string) (string, error)
}

type Adder interface {
	AddPhrase(string, string) error
}

var ErrNotFound = errors.New("err_phrase_not_found")
var ErrPhraseAlreadyExists = errors.New("err_phrase_already_exists")
