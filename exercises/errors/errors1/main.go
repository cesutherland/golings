// errors1
// Make me compile!

package main

import "fmt"

// make this function return an error when b is 0: cannot divide by zero
func divide(a, b float64) (float64, ) {  // a return is missing here
	if b == 0 {
		return 0,  // return your error here
	}

	return a / b, nil
}

func main() {
	ans, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ans)
}
