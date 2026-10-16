package interfacedemo

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {

	return x, nil
}

func ExerciseError() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
