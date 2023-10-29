package arrays

import "fmt"

func Demo3() {
	//finding length of arrays

	numberss := [5]int{8, 50, 90, 4, 1}

	for i := 0; i < len(numberss); i++ {
		fmt.Print(numberss[i])
	}

	//finding the biggest number in array

	for i := 0; i < len(numberss); i++ {
		fmt.Println(numberss[i])
	}

	biggest := numberss[0]

	for i := 0; i < len(numberss); i++ {
		if numberss[i] > biggest {
			biggest = numberss[i]
		} else {
			continue
		}

		fmt.Println(biggest)
	}
}
