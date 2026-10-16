package methoddemo

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y *= f
}

func Pointer() {
	v := Vertex2{3, 4}
	// v.Scale(20)
	fmt.Println(v.Abs())
	fmt.Println(v)

	// 函数需要手动，改成指针类型
	Scale(&v, 10)
	fmt.Println(v)
}
