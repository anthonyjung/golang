package main

import (
	"fmt"
	"math"
)

type MyFloat64 float64

func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat64(-math.Sqrt2)
	fmt.Println(f.Abs())
}
