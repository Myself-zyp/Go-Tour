package flowcontroldemo

import "fmt"

func For() {
	sum := 0
	// for i := 0; i < 10; i++ {
	for i := range 10 {
		sum += i
	}
	fmt.Printf("sum: %v\n", sum)

	result := 1
	// 初始化语句和后置语句都是可选的； while 写法也是这样
	for result < 100 {
		result += result
	}

	fmt.Println(result)

	// 无限循环
	for {
		fmt.Println("forever")
	}

}
