package main

import (
	"errors"
	"fmt"
)

func intDiv(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}

func main() {
	result, remainder, err := intDiv(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)
	}
}
