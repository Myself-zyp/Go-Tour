package moretypesdemo

import "fmt"

func Replace() {
	primes := [3]int{1, 2, 3}
	s := primes[:]
	fmt.Println(s, primes, len(primes), cap(primes))

	s = append(s, 0)

	s1 := primes[:1]
	s1 = append(s1, 12) // 会直接修改数组的第二个值
	fmt.Println(s1)

	s[0] = 11

	fmt.Println(s, primes, len(primes), cap(primes))
}

func Append() {
	var s []int
	PrintSlice(s)

	// append 用于 nil slice
	s = append(s, 0)
	PrintSlice(s)

	// slice 按需增长
	s = append(s, 1)
	PrintSlice(s)

	// 一次添加多个元素
	s = append(s, 2, 3, 4)
	PrintSlice(s)

}
