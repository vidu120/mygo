package main

import (
	"fmt"
)

func takeInputs(acc, initialVel, initialDisp *float64) {
	fmt.Print("\t\tProgram for computing displacement at time t sec..\n")
	fmt.Print("Please input the acceleration first : ")
	fmt.Scan(acc)
	fmt.Print("Please input the initial velocity: ")
	fmt.Scan(initialVel)
	fmt.Print("Please input the inital displacement: ")
	fmt.Scan(initialDisp)
	// fmt.Println(acc, initialVel, initialDisp)
}

//function returning a function
func genDisplacefn(acc, initialVel, initialDisp *float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*(*acc)*t*t + (*initialVel)*t + (*initialDisp)
	}
}

func main() {
	var acc, initialVel, initialDisp float64

	//take inputs from the user
	takeInputs(&acc, &initialVel, &initialDisp)
	// fmt.Println(acc, initialVel, initialDisp)

	dispAtTime := genDisplacefn(&acc, &initialVel, &initialDisp)

	//new line
	fmt.Println()

	//taking time input from user
	fmt.Print("Please input the time for displacement cal.: ")
	var time float64
	fmt.Scan(&time)

	fmt.Printf("Displacement in %.2f sec : %.2f\n", time, dispAtTime(time))

}
