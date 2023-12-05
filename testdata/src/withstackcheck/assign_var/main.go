package assign_to_var

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func externalError() error {
	// var declaration
	//var err error
	//
	//_, err = json.Marshal(nil)
	//if err != nil {
	//	return err // want `error without stacktrace returned from external package`
	//}
	//
	//_, err = json.Marshal(nil)
	//if err != nil {
	//	return errors.WithStack(err)
	//}

	// short declaration
	if _, err := json.Marshal(nil); err != nil {
		return err // want `error without stacktrace returned from external package`
	}

	if _, err := json.Marshal(nil); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func internalError() error {
	// short declaration
	if err := throw(); err != nil {
		return err
	}

	if err := throw(); err != nil {
		return errors.WithStack(err) // want `error with stacktrace returned from internal package`
	}

	return nil
}

func throw() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}
