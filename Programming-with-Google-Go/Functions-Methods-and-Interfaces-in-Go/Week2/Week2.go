package main

import "fmt"

func main() {

	fmt.Printf("Enter values for acceleration, initial velocity, and initial displacement:  ")
	var a, v0, s0 float64
	fmt.Scan(&a, &v0, &s0)

	fmt.Printf("Enter a value for time: ")
	var t float64
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Total Displacement is: %f", fn(t))


}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(time float64) float64 {
		return (0.5 * a * (time * time)) + (v0*time) + (s0)
	}
}


