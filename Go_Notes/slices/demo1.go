package slices

import "fmt"

func Demo1() {
	/*The "slice" structure allows us to use the array structure more dynamically.
	It allows us to make additions and subtractions easily.
	*/
	names := make([]string, 2)
	names[0] = "Bora"
	names[1] = "Burak"

	fmt.Println(names)
	fmt.Println(len(names))

	names = append(names, "Tahsin")

	fmt.Println(names)
	fmt.Println(len(names))
}
