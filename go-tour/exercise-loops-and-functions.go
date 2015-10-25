/*
 * exercise-loops-and-functions.go
 *   personal practise by Tao Ma
 */
package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
)

func Sqrt10times(x float64) float64 {
	var z float64 = 1

	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}

	return z
}

func SqrtStable(x float64) float64 {
	var z float64 = 1
	var lastz float64 = z

	for {
		lastz = z
		z = z - (z*z-x)/(2*z)
		// Is the following case valid ?
		if lastz == z {
			break
		}
	}

	return z
}

func SqrtDelta(x float64, delta float64) float64 {
	var z float64 = 1
	var lastz float64 = z

	for {
		lastz = z
		z = z - (z*z-x)/(2*z)
		if math.Abs(lastz-z) < delta {
			break
		}
	}

	return z
}

func main() {
	var times int
	var mode string
	var delta float64
	var result float64

	flag.IntVar(&times, "times", 1, "run specified times")
	flag.StringVar(&mode, "mode", "n10", "method of computation of square root.")
	flag.Float64Var(&delta, "delta", 1e-5, "control the delta between Newton's method and standard Math.Sqrt.")

	flag.Parse()

	if flag.NArg() != 1 {
		errors.New("need a number")
	}

	v, err := strconv.ParseFloat(flag.Arg(0), 64)
	if err != nil {
		errors.New("invalid float number")
	}

	switch mode {
	case "std":
		result = math.Sqrt(v)
		for ; times > 0; times-- {
			_ = math.Sqrt(v)
		}
	case "n10":
		result = Sqrt10times(v)
		for ; times > 0; times-- {
			_ = Sqrt10times(v)
		}
	case "ndelta":
		result = SqrtDelta(v, delta)
		for ; times > 0; times-- {
			_ = SqrtDelta(v, delta)
		}
	case "nstable":
		result = SqrtStable(v)
		for ; times > 0; times-- {
			_ = SqrtStable(v)
		}
	default:
		errors.New("invalid mode")
	}

	fmt.Println(result)
}
