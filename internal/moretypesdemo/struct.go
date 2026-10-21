package moretypesdemo

import "fmt"

func Struct() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v)

	p := &v
	p.X = 1e9 // 1 乘 10 的 9 次方
	// (*p).X 的简写方式，无需显式解引用
	fmt.Println(v)

}
