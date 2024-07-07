package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) milesleft() uint8 {
	return e.mpg * e.gallons
}

type engine interface {
	milesleft() uint8
}

func main() {
	enginer := gasEngine{
		mpg:     0,
		gallons: 0,
	}
	fmt.Println(enginer.mpg)
}
