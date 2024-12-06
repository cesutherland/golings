// errors2
// Make me compile!

// I AM NOT DONE
package main

import (
	"errors"
	"fmt"
	"testing"
)

type ErrDivideByZero struct { // build the ErrDivideByZero struct

}

func (e ErrDivideByZero)  string { // implement the error interface
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
