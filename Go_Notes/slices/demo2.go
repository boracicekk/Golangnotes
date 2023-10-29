package slices

import "fmt"

func Demo2() {
	cities := []string{"Istanbul,Ankara,Antalya"}
	fmt.Println(cities)
	citiescopy := make([]string, 3)
	fmt.Println(citiescopy)
	copy(citiescopy, cities)
	fmt.Println(citiescopy)

	cities = append(cities, "Izmir")
	fmt.Println(cities)
	fmt.Println(citiescopy)
	//As you can see, we added an element to the "cities" array, but the "citiescopy" array was not affected.

	///Now, we'll learn to filtering process.///

	//This statement show us, 1st index to 3rd but NOT INCLUDE 3rd element.
	fmt.Println(cities[1:3])

	//This statement show us, 1st index to last element printing.
	fmt.Println(cities[1:])

	//This statement show us, first element to 3th element printing But, 3th element not include.
	fmt.Println(cities[:3])

}
