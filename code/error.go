package code

import (
	"fmt"

	"golang.org/x/xerrors"
)

type privateError struct {
	code Code
	err  error
}

func (e privateError) Error() string {
	return fmt.Sprintf("code: %s, error: %s", e.code, e.err)
}

func Error(c Code, format string) error {
	if c == OK {
		return nil
	}
	return privateError{code: c, err: xerrors.New(format)}
}
