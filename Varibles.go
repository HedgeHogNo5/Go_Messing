package main

import (
	"fmt"
)

func Sum() {
	var x int
	x = 5         // this is the first way you can create a varible. It is defined as an integer
	var y int = 7 // this is teh second way you can make an
	z := 9        // this is the final way you can create a variable,
	sum := x + y + z
	fmt.Println(sum)
}

func Avrg_3(x float64, y float64, z float64) float64 {
	return(x + y + z) / 3 //This takes in the float values of x,y,z and will output the average of the three 
}

func Avrg_n() {
	
}