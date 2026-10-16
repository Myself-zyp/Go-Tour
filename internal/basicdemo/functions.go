package basicdemo

import "fmt"

// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

// 多个返回值
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
func split(sum int) (x, y int) { // 命名返回值会被视为定义在函数顶部的变量
	x = sum * 4 / 9
	y = sum - x
	return // 裸返回，只适合在较短的函数中使用，长函数中影响可读性
}

func Fn() {
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
