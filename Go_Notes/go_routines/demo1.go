package go_routines

import (
	"fmt"
	"time"
)

// I learn to "go" struct. The "Go" structure allows the desired commands to run synchronously.
func EvenNumbers() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("Even numbers: ", i)
		time.Sleep(1 * time.Second)
	}
}

func OddNumbers() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("Odd numbers: ", i)
		time.Sleep(1 * time.Second)
	}
}

//I learned to run 2 different functions with the help of "go routine".
