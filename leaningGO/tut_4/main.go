package main

import "fmt"

func main() {
	var intArr [3]int
	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3
	fmt.Println(intArr)

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//slices

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)

	mymap := map[string]uint8{"one": 1, "two": 2, "three": 3}
	fmt.Println(mymap)

	for i, v := range intArr {
		fmt.Printf("Index : %v, Value: %v \n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
