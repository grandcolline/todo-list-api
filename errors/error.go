package errors

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/todo-list-api/errors/code"
	"google.golang.org/grpc/codes"
)

type privateError struct {
	code codes.Code
	err  error
}

func (e privateError) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s", e.code, e.err)
}

// Errorf returns an error containing an error code and a description;
// Errorf returns nil if c is OK.
func Errorf(c code.Code, format string, a ...interface{}) error {
	if c == code.OK {
		return nil
	}
	return privateError{
		code: c,
		err:  errors.Errorf(format, a...), // github.com/pkg/errorsでラップする
	}
}

// Code returns the error code for err if it was produced by this system.
// Otherwise, it returns codes.Unknown.
func Code(err error) code.Code {
	if err == nil {
		return code.OK
	}
	if e, ok := err.(privateError); ok {
		return e.code
	}
	return code.Unknown
}

// StackTrace shows stacktrace. If error is not private error, this returns empty string.
func StackTrace(err error) string {
	if e, ok := err.(privateError); ok {
		return fmt.Sprintf("%+v\n", e.err)
	}
	return ""
}
