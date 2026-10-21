package moretypesdemo

import (
	"fmt"
	"strings"
)

func MakeSlice() {
	a := make([]int, 5)
	fmt.Println(len(a), cap(a))

	b := make([]int, 2, 6)
	fmt.Println(len(b), cap(b))
	PrintSlice(b)

	c := b[:2]
	PrintSlice(c)

	d := c[2:5]
	PrintSlice(d)

	// slice of slices 切片的切片
	board := [][]string{
		// []string{"_", "_", "_"},
		// []string{"_", "_", "_"},
		// []string{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	fmt.Println(board)
	for i := range len(board) {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
