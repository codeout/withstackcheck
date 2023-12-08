package named_return

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func externalError() (err error) {
	// var declaration
	_, err = json.Marshal(nil)
	if err != nil {
		return // want `error without stacktrace returned from external package`
	}

	return nil
}

func externalErrorWithOtherReturns() (other1, err, other2 error) {
	// var declaration
	_, err = json.Marshal(nil)
	if err != nil {
		return // want `error without stacktrace returned from external package`
	}

	return nil, nil, nil
}

func internalError() (err error) {
	// var declaration
	err = throw()
	if err != nil {
		return
	}

	return nil
}

func internalErrorWithOtherReturns() (other1, err, other2 error) {
	// var declaration
	err = throw()
	if err != nil {
		return
	}

	return nil, nil, nil
}

func throw() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}
