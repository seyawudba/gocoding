package main

import (
	"fmt"
	"math"
)

func main() {
	params := map[string]float64{"Acceleration": 0, "Initial Velocity": 0, "Initial Displacement": 0}

	for equation_key, _ := range params {
		fmt.Printf("Please enter value for %s:  ", equation_key)
		var entry float64
		_, err := fmt.Scan(&entry)
		if err != nil {
			return
		}
		params[equation_key] = entry
	}

	a := params["Acceleration"]
	v := params["Initial Velocity"]
	s := params["Initial Displacement"]

	fn := GenDisplaceFn(a, v, s)

	fmt.Println(fn)

	fmt.Printf("Please enter the time value:  ")
	var t float64
	_, err := fmt.Scan(&t)
	if err != nil {
		return
	}

	fmt.Println(fn(t))

}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {

	fn := func(t float64) float64 {

		fmt.Println(0.5*a*math.Pow(t, 2) + v*t + s)

		return 0.5*a*math.Pow(t, 2) + v*t + s

	}

	return fn

}
