package errors

import (
	"fmt"
	"strings"
)

// ServiceError descripes a behaviour of the service error.
type ServiceError interface {
	Code() int
	Text() string
	Error() string
	Method(method string) *Error
	Scope(key string, value interface{}) *Error
}

// Error represents a service level error.
type Error struct {
	code   int
	text   string
	method string
	scope  string
	err    error
}

// New returns the new Error.
func New(text string) *Error {
	return &Error{
		text: text,
	}
}

// NewWithCode returns the new Error with a code.
func NewWithCode(code int, text string) *Error {
	return &Error{
		code: code,
		text: text,
	}
}

// clone returs a clone of the current error.
func (e *Error) clone() *Error {
	return &Error{
		code:  e.code,
		text:  e.text,
		err:   e.err,
		scope: e.scope,
	}
}

// Wrap wraps a error.
func (e *Error) Wrap(err error) *Error {
	// if the previous error has the same error code
	// do not wrap the error
	preverr, ok := err.(*Error)
	if ok && preverr.Code() == e.Code() {
		return preverr
	}

	// wrap the error
	out := e.clone()

	out.err = err

	return out
}

// Unwrap returns a text of the error.
func (e *Error) Unwrap() error {
	return e.err
}

// Error returns a service level error message.
func (e *Error) Error() (out string) {
	if e.code != 0 {
		out = fmt.Sprintf("%d: ", e.code)
	}

	out = fmt.Sprintf("%s%s", out, e.Text())

	return out
}

// Method add a method name to the error text.
func (e *Error) Method(method string) *Error {
	out := e.clone()

	out.method = method

	return out
}

// Scope add a scope to the error.
func (e *Error) Scope(key string, value interface{}) *Error {
	out := e.clone()

	out.scope = strings.Trim(fmt.Sprintf("%s=%s %v", key, value, out.scope), " ")

	return out
}

// Code returns a code of the error.
func (e *Error) Code() int {
	return e.code
}

// Text returns a text of the error.
func (e *Error) Text() (out string) {
	out = fmt.Sprintf("%s%s", out, e.text)

	if e.method != "" {
		out = fmt.Sprintf("%s: %s", out, e.method)
	}

	if e.scope != "" {
		out = fmt.Sprintf("%s: %s", out, e.scope)
	}

	if e.err != nil {
		out = fmt.Sprintf("%s <-- %s", out, e.err)
	}

	return out
}
