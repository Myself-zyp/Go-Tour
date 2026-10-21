package moretypesdemo

import (
	"fmt"
	"strings"
)

type vertex struct {
	lat, long float64
}

func basic() {
	m := make(map[string]vertex)
	m["Bell Labs"] = vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m)
}

func literals() {
	// literals
	m1 := map[string]vertex{
		"Bell":   {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m1)
}

func mutating() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	result := make(map[string]int)

	// for _, word := range strings.Fields(s) {
	for word := range strings.FieldsSeq(s) { // 新方法
		fmt.Println(string(word))
		result[word]++
	}

	return result
}

func Map() {
	// mutating()
	fmt.Println(WordCount("hello web go lang go"))
}
