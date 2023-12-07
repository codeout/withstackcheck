package _func

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// single return
func externalError() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}

func externalErrorWithStack() error {
	return errors.WithStack(errors.New("error"))
}

func internalError() error {
	return throw1()
}

func internalErrorWithStack() error {
	return errors.WithStack(throw1()) // want `error with stacktrace returned from internal package`
}

// multiple return
func externalErrorWithMultiReturn() ([]byte, error) {
	return json.Marshal(nil) // want `error without stacktrace returned from external package`
}

func internalErrorWithMultiReturn() ([]byte, error) {
	return throw2()
}

// method
type structUnderTest1 struct {
	sut structUnderTest2
}
type structUnderTest2 struct{}

func (_ structUnderTest2) throw() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}

func internalErrorFromMethod() error {
	return structUnderTest1{}.sut.throw()
}

func internalErrorFromMethodWithStack() error {
	return errors.WithStack(structUnderTest1{}.sut.throw()) // want `error with stacktrace returned from internal package`
}

// interface
type interfaceUnderTest interface {
	throw() error
}
type structUnderTest3 struct {
	iut interfaceUnderTest
}

func internalErrorFromInterface(iut interfaceUnderTest) error {
	return iut.throw()
}

func internalErrorFromInterfaceWithStack(iut interfaceUnderTest) error {
	return errors.WithStack(iut.throw()) // want `error with stacktrace returned from internal package`
}

func internalErrorFromInterfaceViaMethod(sut structUnderTest3) error {
	return sut.iut.throw()
}

func internalErrorFromInterfaceViaMethodWithStack(sut structUnderTest3) error {
	return errors.WithStack(sut.iut.throw()) // want `error with stacktrace returned from internal package`
}

func throw1() error {
	return errors.New("error") // want `error without stacktrace returned from external package`
}

func throw2() ([]byte, error) {
	return json.Marshal(nil) // want `error without stacktrace returned from external package`
}
