package arrays

import "fmt"

func Demo2() {

	//How to define an array's element by one by or one time?
	var numberss [5]int
	numberss[0] = 1
	numberss[1] = 2
	numberss[2] = 3
	numberss[3] = 4
	numberss[4] = 5
	count := 5

	for i := 0; i < count; i++ {
		fmt.Print(numberss[i])
	}

	numbersss := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < count; i++ {
		fmt.Print(numbersss[i])
	}

}
