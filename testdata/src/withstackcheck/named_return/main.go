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

	// short declaration
	if _, err := json.Marshal(nil); err != nil {
		return // want `error without stacktrace returned from external package`
	}

	return nil
}

func internalError() (err error) {
	// var declaration
	err = throw()
	if err != nil {
		return
	}

	// short declaration
	if err := throw(); err != nil {
		return
	}

	return nil
}

func throw() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}
