package arrays

import "fmt"

func Demo4() {
	// Im learning to multiple dimension arrays.

	var numbers [2][3]int
	numbers[0][0] = 1
	numbers[0][1] = 3
	numbers[0][2] = 5
	numbers[1][0] = 2
	numbers[1][1] = 4
	numbers[1][2] = 6

	/* if this comment's bottom will active you can see the error message "error bounds".

	numbers[2][0] = 6

	or

	numbers[0][3] = 6

	 When we mention the arrays' index number, we should realize this is a index number not capacity of array.
	*/

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(numbers[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
