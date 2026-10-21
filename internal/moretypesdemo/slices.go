package moretypesdemo

import "fmt"

func Slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)

	// 切片默认边界
	// var a [10]int
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]
	// 这些 slice 表达式都是等价的
	fmt.Println(q[1:4], q[:2], q[1:])
}
