package range1

import "fmt"

//A function example function that adds odd numbers between 1-10
func Demo2() {
	OnetoTen := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	add := 0
	for _, num := range OnetoTen {
		if num%2 != 0 {
			add += num
		}

	}
	fmt.Println(add)
}
