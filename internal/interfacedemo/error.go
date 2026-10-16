package interfacedemo

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", "time", e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func ErrorFn() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
