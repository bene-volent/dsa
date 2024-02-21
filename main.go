package main

import (
	"fmt"

	"github.com/bene-volent/dsa/array"
)

func main() {
	var arr, arr1 = array.New[int](), array.New[int]()

	for i := range 80 {
		arr.PushElement(i)
	}
	for j := range 21 {
		arr1.PushElement(j)

	}
	arr.PrintAll()
	arr1.PrintAll()
	var mergedArr, err = arr.Merge(&arr1)
	mergedArr.PrintAll()
	fmt.Println(err)

}
