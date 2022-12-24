package main

import (
	"fmt"
)

func sum() {
	var x int
	x = 5         // this is the first way you can create a varible. It is defined as an integer
	var y int = 7 // this is teh second way you can make an
	z := 9        // this is the final way you can create a variable,
	sum := x + y + z
	fmt.Println(sum)
	//a varible can also be defined in the code i.e. float64(var) would make var a float.
}

func avrgThree(x float64, y float64, z float64) float64 {
	if x+y+z == 0 {
		return 0
	}
	return (x + y + z) / 3 //This takes in the float values of x,y,z and will output the average of the three
}

func avrgN(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range xs {
		sum += v
	}
	return sum / float64(len(xs))
}

//This average takes in an array and outputs the average
//If the array is empty it outputs 0
