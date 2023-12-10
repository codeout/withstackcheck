package type_assert

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func externalError() error {
	// var declaration
	var err error

	_, err = json.Marshal(nil)
	if err != nil {
		return err.(error) // want `error without stacktrace returned from external package`
	}

	_, err = json.Marshal(nil)
	if err != nil {
		return errors.WithStack(err.(error))
	}

	// short declaration
	if _, err := json.Marshal(nil); err != nil {
		return err.(error) // want `error without stacktrace returned from external package`
	}

	if _, err := json.Marshal(nil); err != nil {
		return errors.WithStack(err.(error))
	}

	return nil
}

func internalError() error {
	// var declaration
	var err error

	err = throw()
	if err != nil {
		return err.(error)
	}

	err = throw()
	if err != nil {
		return errors.WithStack(err.(error)) // want `error with stacktrace returned from internal package`
	}

	// short declaration
	if err := throw(); err != nil {
		return err.(error)
	}

	if err := throw(); err != nil {
		return errors.WithStack(err.(error)) // want `error with stacktrace returned from internal package`
	}

	return nil
}

func throw() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}
