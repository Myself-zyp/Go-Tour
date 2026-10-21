package flowcontroldemo

import "fmt"

func Defer() {
	x := 10
	defer fmt.Println(x)

	defer func() { fmt.Println(x) }()

	x = 20
	fmt.Println("正常执行：", x)

	fmt.Println("counting")

	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// defer fmt.Println("world")

	// fmt.Println("hello")
}
