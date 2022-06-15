package main

import "fmt"

//const p = 500
//const t = .78

func p1(d float64) {

	var p float64 = 500
	var t float64 = .78
	fmt.Println("Paycheck: ", (p*d)*t)
}

func main() {

	p1(10)

}

//
