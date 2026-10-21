package basicdemo

import "fmt"

const (
	// 将 1 左移 100 位， 得到一个非常大的数
	Big = 1 << 100

	// 再将它向右移 99 位，最终得到 1 << 1, 也就是 2
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func NumericConstants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))
}
