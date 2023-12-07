package other_package

import (
	"withstackcheck/other_package/internal"

	"github.com/pkg/errors"
)

func internalError() error {
	// var declaration
	var err error

	err = internal.Throw()
	if err != nil {
		return err
	}

	err = internal.Throw()
	if err != nil {
		return errors.WithStack(err) // want `error with stacktrace returned from internal package`
	}

	// short declaration
	if err := internal.Throw(); err != nil {
		return err
	}

	if err := internal.Throw(); err != nil {
		return errors.WithStack(err) // want `error with stacktrace returned from internal package`
	}

	return nil
}

func internalVarError() error {
	return internal.VarError
}

func internalVarErrorWithStack() error {
	return errors.WithStack(internal.VarError) // want `error with stacktrace returned from internal package`
}
