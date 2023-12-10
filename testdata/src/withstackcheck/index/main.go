package assign_to_var

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func externalError() error {
	var errs []error

	if _, err := json.Marshal(nil); err != nil {
		errs = append(errs, err)
		return errs[0] // want `error without stacktrace returned from external package`
	}

	if _, err := json.Marshal(nil); err != nil {
		errs = append(errs, errors.WithStack(err))
		return errs[0]
	}

	return nil
}

func internalError() error {
	var errs []error

	if err := throw(); err != nil {
		errs = append(errs, err)
		return err
	}

	if err := throw(); err != nil {
		errs = append(errs, errors.WithStack(err))
		return errs[0] // want `error with stacktrace returned from internal package`
	}

	return nil
}

func throw() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}
