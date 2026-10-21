package moretypesdemo

import "fmt"

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 是隐式的
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

func StructLiterals() {
	fmt.Println(v1, p, v2, v3)
}
