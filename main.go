package main

import (
	"fmt"
)

const p float64 = 500
const t float64 = .78

//const d = 10

func p1(d float64) {
	//var p float64 = 500
	//var t float64 = .78

	fmt.Println("Paycheck: ", (p*d)*t)
	fmt.Printf("What I make a month: %v if I work %v days\n", (p*d)*t*2, d*2)

}

func main() {

	// func p1(d, float64) int{
	// 	var p float64 = 500
	// 	var t float64 = .78

	// 	fmt.Println("Paycheck: ", (p*d)*t)
	// }

	p1(11)
	//fmt.Println("Make a month", p1*m)

	//var p float64 = 500
	//var t float64 = .78
	//fmt.Println("Paycheck: ", (p*d)*t)

}
