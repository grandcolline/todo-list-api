package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
)

type codeError struct {
	code  Code
	err   error
	cause string
}

func New(c Code, msg string) error {
	if c == OK {
		return nil
	}
	return &codeError{
		code:  c,
		err:   errors.New(msg),
		cause: getCause(2),
	}
}

func AddCode(c Code, format string, a ...interface{}) error {
	if c == OK {
		return nil
	}
	return &codeError{
		code:  c,
		err:   fmt.Errorf(format, a...),
		cause: getCause(2),
	}
}

func Errorf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	return &codeError{
		code:  GetCode(err),
		err:   err,
		cause: getCause(2),
	}
}

func (e codeError) Error() string {
	return e.err.Error() + " (" + e.cause + ")"
}

func Format(e error) string {
	return fmt.Sprintf("Code: %s, Msg: %s", GetCode(e), e.Error())
}

func GetCode(e error) Code {
	if e == nil {
		return OK
	}
	c := &codeError{}
	if ok := errors.As(e, &c); ok {
		return c.code
	}
	return Unknown
}

func getCause(skip int) string {
	if _, filename, line, ok := runtime.Caller(skip); ok {
		return filename + ":" + strconv.Itoa(line)
	}
	return "unknown"
}
