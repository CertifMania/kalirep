package main

import (
	"fmt"
	"math"
)

func main() {
	// Max and Min values for int32
	fmt.Printf("\nint32:\t Min = %d, \t\tMax = %d\n", math.MinInt32, math.MaxInt32)

	// Max and Min values for int
	fmt.Printf("int:\t Min = %d, \tMax = %d\n", math.MinInt, math.MaxInt)

	// Max value for uint32
	fmt.Printf("uint32:\t Max = %d\n", math.MaxUint32)

	// Max value for uint
	fmt.Printf("uint:\t Max = %d\n", ^uint(0))

	// Explanation:
	// math.MinInt32, math.MaxInt32: Predefined constants for int32 in math package.
	// math.MinInt, math.MaxInt: Predefined constants for int in Go.
	// math.MaxUint32: Predefined constant for uint32 in math package.
	// ^uint(0): Bitwise NOT of 0 gives the maximum value for uint.
}
