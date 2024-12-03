// errors2
// Make me compile!

package main

import (
	"errors"
	"fmt"
)

type ErrDivideByZero struct {  // build the ErrDivideByZero struct
}

func (e ErrDivideByZero)  string {  // implement the error interface
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

func main() {

	ans, err := divide(10, 0)

	if err != nil {
		var e ErrDivideByZero
		switch {
		case errors.As(err, &e):
			fmt.Printf("%.2f / %.2f is not mathematically valid: %s\n",
				e.Number1, e.Number2, e.Error())
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
	}

	fmt.Println(ans)
}
