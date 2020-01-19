package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"

	"github.com/grandcolline/todo-list-api/util/errors/code"
)

type codeError struct {
	code  code.Code
	err   error
	cause string
}

func New(c code.Code, msg string) error {
	if c == code.OK {
		return nil
	}
	return &codeError{
		code:  c,
		err:   errors.New(msg),
		cause: getCause(2),
	}
}

func NewFromFmt(m errMsg, a ...interface{}) error {
	if m.Code == code.OK {
		return nil
	}
	return &codeError{
		code:  m.Code,
		err:   fmt.Errorf(m.Msg, a...),
		cause: getCause(2),
	}
}

func AddCode(c code.Code, format string, a ...interface{}) error {
	if c == code.OK {
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

func GetCode(e error) code.Code {
	if e == nil {
		return code.OK
	}
	c := &codeError{}
	if ok := errors.As(e, &c); ok {
		return c.code
	}
	return code.Unknown
}

func getCause(skip int) string {
	if _, filename, line, ok := runtime.Caller(skip); ok {
		return filename + ":" + strconv.Itoa(line)
	}
	return "unknown"
}
