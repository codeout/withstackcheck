package reassign_var

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func varDeclaration() error {
	var err error

	_, err = json.Marshal(nil)
	if err != nil {
		return err // want `error without stacktrace returned from external package`
	}

	err = throw()
	if err != nil {
		return err
	}

	_, err = json.Marshal(nil)
	if err != nil {
		return errors.WithStack(err)
	}

	err = throw()
	if err != nil {
		return errors.WithStack(err) // want `error with stacktrace returned from internal package`
	}

	_, err = json.Marshal(nil)
	if err != nil {
		return err // want `error without stacktrace returned from external package`
	}

	return nil
}

func shortDeclaration() error {
	_, err := json.Marshal(nil)
	if err != nil {
		return err // want `error without stacktrace returned from external package`
	}

	err = throw()
	if err != nil {
		return err
	}

	_, err = json.Marshal(nil)
	if err != nil {
		return errors.WithStack(err)
	}

	err = throw()
	if err != nil {
		return errors.WithStack(err) // want `error with stacktrace returned from internal package`
	}

	_, err = json.Marshal(nil)
	if err != nil {
		return err // want `error without stacktrace returned from external package`
	}

	return nil
}

func throw() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}
