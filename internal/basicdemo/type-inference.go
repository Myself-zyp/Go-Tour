package basicdemo

import "fmt"

func TypeInference() {
	v := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)
}
