package mapdemo

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m2 = map[string]Vertex{
	"Bell Labs": {40.68, -74.39},
	"Google":    {36.22, -122.08},
}

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	
	for _, word := range strings.Fields(s) {
		result[word]++
	}
	
	return result
}


func Run() {
	// map 练习
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m2)

	m3 := make(map[string]int)
	m3["Answer"] = 30
	fmt.Println("The value: ", m3["Answer"])
	m3["Answer"] = 60
	fmt.Println("The value: ", m3["Answer"])


	v, ok := m3["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)
	
	delete(m3, "Answer")
	fmt.Println("The value: ", m3["Answer"])

	v, ok = m3["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)

	fmt.Println("Return", WordCount("word count"))

	
}