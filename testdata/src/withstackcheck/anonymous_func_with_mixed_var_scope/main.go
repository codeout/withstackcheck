package anonymous_func_with_mixed_var_scope

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func externalError() error {
	var f func() error

	// var declaration
	var err error

	f = func() error {
		_, err = json.Marshal(nil)
		if err != nil {
			return err // want `error without stacktrace returned from external package`
		}

		return nil
	}

	f = func() error {
		_, err = json.Marshal(nil)
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}

	return f()
}

func internalError() error {
	var f func() error

	// var declaration
	var err error

	f = func() error {
		err = throw()
		if err != nil {
			return err
		}

		return nil
	}

	f = func() error {
		err = throw()
		if err != nil {
			return errors.WithStack(err) // want `error with stacktrace returned from internal package`
		}

		return nil
	}

	return f()
}

func throw() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}
