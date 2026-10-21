package flowcontroldemo

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)

	for range 10 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func Exercise() {
	fmt.Println(Sqrt(2))
}
