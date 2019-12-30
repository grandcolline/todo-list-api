package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type customError struct {
	code Code
	err  error
}

// New returns an error code and error witr the supplied message.
// New returns nil if c is OK.
func New(c Code, msg string) error {
	if c == OK {
		return nil
	}
	return customError{
		code: c,
		err:  errors.New(msg),
	}
}

// Errorf returns an error containing an error code and a description;
// Errorf returns nil if c is OK.
func Errorf(c Code, format string, a ...interface{}) error {
	if c == OK {
		return nil
	}
	return customError{
		code: c,
		err:  errors.Errorf(format, a...),
	}
}

func (e customError) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s", e.code, e.err)
}

// GetCode retirn error code.
func GetCode(err error) Code {
	if err == nil {
		return OK
	}
	if e, ok := err.(customError); ok {
		return e.code
	}
	return Unknown
}

// StackTrace shows stacktrace. If error is not private error, this returns empty string.
func StackTrace(err error) string {
	if e, ok := err.(customError); ok {
		return fmt.Sprintf("%+v\n", e.err)
	}
	return ""
}
