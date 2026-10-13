package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m2 = map[string]Vertex{
	"Bell Labs": {40.68, -74.39},
	"Google":    {36.22, -122.08},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m2)
}
