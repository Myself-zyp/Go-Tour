package moretypesdemo

import "fmt"

func Range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	po := make([]int, 10)
	for i := range po {
		po[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println(po)

	for _, value := range po {
		fmt.Printf("%d\n", value)
	}
}
