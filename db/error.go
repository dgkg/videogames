package db

import "fmt"

type Error struct {
	errType string
	err     error
	params  map[string]interface{}
}

const (
	ErrDBQuery     = "query"
	ErrDBConnexion = "connexion"
)

func (e Error) Error() string {
	return fmt.Sprintf("type %v message %v params %v", e.errType, e.err, e.params)
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func NewQueryErr(map[string]interface{}) error {
	return &Error{
		errType: ErrDBQuery,
	}
}
