// errors2
// Make the tests pass!

package main

import (
	"errors"
	"fmt"
	"testing"
)

type ErrDivideByZero struct { // build the ErrDivideByZero struct
	Number1 float64
	Number2 float64
	Message string
}

func (e ErrDivideByZero) Error() string { // implement the error interface
	return fmt.Sprintf("divide by zero error (%s)", e.Message)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &ErrDivideByZero{
			Number1: a,
			Number2: b,
			Message: fmt.Sprintf("cannot divide by zero"),
		}
	}

	return a / b, nil
}

func TestErrDivideByZero(t *testing.T) {
	_, err := divide(10, 0)

	var e *ErrDivideByZero

	if !errors.As(err, &e) {
		t.Error("e should be an ErrDivideByZero error")
	}
}
