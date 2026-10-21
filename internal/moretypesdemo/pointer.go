package moretypesdemo

import "fmt"

func Pointer() {
	var p *int
	fmt.Println(p)

	i := 42
	p = &i
	fmt.Println(i, p, *p)

	m, n := 42, 2701
	p1 := &m         // 指向 m
	fmt.Println(*p1) // 通过指针读取 m
	*p1 = 21         // 通过指针设置 m
	fmt.Println(m)   // 查看 m 的新值

	p1 = &n        // 指向 n
	*p1 = *p1 / 37 // 通过指针对 n 进行除法运算
	fmt.Println(n) // 查看 n 的新值

}
