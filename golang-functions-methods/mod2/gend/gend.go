package main

import (
	"fmt"
	"log"
)

// GenDisplacement generates a displacement function, for given parameters
func GenDisplacement(a, v0, s0 float64) func(float64) float64 {
	ha := a / 2.0
	return func(t float64) float64 {
		return ha*t*t + v0*t + s0
	}
}

func main() {
	var a, v0, s0 float64
	fmt.Print("Acceleration: ")
	if _, err := fmt.Scan(&a); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Initial velocity: ")
	if _, err := fmt.Scan(&v0); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Initial displacement: ")
	if _, err := fmt.Scan(&s0); err != nil {
		log.Fatal(err)
	}

	fn := GenDisplacement(a, v0, s0)

	for _, t := range []float64{3, 5} {
		fmt.Printf("t = %.0f -> s = %f\n", t, fn(t))
	}
}
