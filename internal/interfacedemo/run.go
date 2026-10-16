package interfacedemo

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法意味着类型 T 实现了，接口 I
// 但无需显示的声明这一点
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("%v, %T\n", i, i)
}

type N struct {
	S string
}

func (n *N) M() {
	if n == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(n.S)
}

func Run() {
	fmt.Println("interface demo")

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // Myfloat 实现了 Abser
	a = &v // *Vertex 实现了 Abser

	// 此时 v 是 Vertex 而不是 *Vertex， 因此它没有实现 Abser
	// a = v

	fmt.Println(a.Abs())

	var i I
	i = &T{"hello"}
	describe(i)
	i.M()

	var f2 F
	// i = F(math.Pi)
	i = f2
	describe(i)
	i.M()

	var n *N
	i = n
	describe(i)
	i.M()

	var i2 I
	describe(i2)

	var i3 interface{}
	describe1(i3)

	i3 = 42
	describe1(i3)
	i3 = "hello"
	describe1(i3)

	Assertions()
	Switches()
	Stringer()
	Exercise()
	ErrorFn()
	ExerciseError()
}

func describe1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
