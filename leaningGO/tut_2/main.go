package main

import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 1
	fmt.Println(intNum)

	var floatNum float64 = 1.0
	fmt.Println(floatNum)

	// cant add between inter types

	text2 := foo()
	fmt.Println(text2)
	fmt.Println(utf8.RuneCountInString("helo world"))
	var mystring string = "helo world"
	fmt.Println(mystring)

	var number int
	fmt.Println(number)
}

func foo() string {
	return "hello world"
}
