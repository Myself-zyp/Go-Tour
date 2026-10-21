package moretypesdemo

import "fmt"

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func SliceLenCap() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	s := primes[:]
	fmt.Println(primes)
	PrintSlice(s)

	// 重新切片 length 0
	s = s[:0]
	PrintSlice(s)

	// 扩展长度
	s = s[:4]
	PrintSlice(s)

	// 去掉前两个元素
	// 切片的容量是从切片的第一个元素开始计算，底层数组中可用的元素个数。
	s = s[2:]
	PrintSlice(s) // cap 4

	var ns []int
	fmt.Println(ns, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil!")
	}
}
