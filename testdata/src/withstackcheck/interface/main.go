package _interface

import (
	"github.com/pkg/errors"
)

type interfaceUnderTest interface {
	throw() error
}

type structUnderTest struct {
	iut interfaceUnderTest
}

func errorFromInterface(iut interfaceUnderTest) error {
	// var declaration
	var err error

	err = iut.throw()
	if err != nil {
		return err
	}

	err = iut.throw()
	if err != nil {
		return errors.WithStack(err) // want `error with stacktrace returned from internal package`
	}

	// short declaration
	if err := iut.throw(); err != nil {
		return err
	}

	if err := iut.throw(); err != nil {
		return errors.WithStack(err) // want `error with stacktrace returned from internal package`
	}

	return nil
}

func errorFromInterfaceViaMethod(sut structUnderTest) error {
	// var declaration
	var err error

	err = sut.iut.throw()
	if err != nil {
		return err
	}

	err = sut.iut.throw()
	if err != nil {
		return errors.WithStack(err) // want `error with stacktrace returned from internal package`
	}

	// short declaration
	if err := sut.iut.throw(); err != nil {
		return err
	}

	if err := sut.iut.throw(); err != nil {
		return errors.WithStack(err) // want `error with stacktrace returned from internal package`
	}

	return nil
}
