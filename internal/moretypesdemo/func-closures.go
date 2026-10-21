package moretypesdemo

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	// 前两个元素之和，保证有前两个元素 0，1
	a, b := 0, 1

	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func FuncClosures() {
	pos, neg := adder(), adder()
	for i := range 10 {
		fmt.Println(pos(i), neg(-1*i))
	}

	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}
