package basicdemo

import "fmt"

var c, python, java bool

// initiallizers
var m, n int = 1, 2

func Variables() {
	var i int
	var javaScript, rust = true, "no!"
	fmt.Println(i, c, python, java, m, n, javaScript, rust)
}
