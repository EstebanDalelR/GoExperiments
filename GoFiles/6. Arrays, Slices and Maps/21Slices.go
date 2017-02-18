package main

import "fmt"

func main() {

	var x []float64
	y := make([]float64, 5, 10)
	//make([]type, slice size, array size)
	//you can cut up an array with x:= arr[0:len(arr)]
	fmt.Println(x)

}
