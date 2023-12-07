package internal

import "errors"

var VarError = errors.New("error")

func Throw() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}
