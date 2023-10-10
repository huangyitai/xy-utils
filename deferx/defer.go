package deferx

import (
	"github.com/pkg/errors"
	"io"
)

// PanicToError ...
func PanicToError(e *error) {
	if i := recover(); i != nil {
		if e == nil {
			return
		}

		if err, ok := i.(error); ok {
			*e = err
		} else {
			*e = errors.Errorf("%+v", i)
		}
	}
}

// Close ...
func Close(closer io.Closer, e *error) {
	err := closer.Close()
	if e != nil && *e == nil {
		*e = err
	}
}
