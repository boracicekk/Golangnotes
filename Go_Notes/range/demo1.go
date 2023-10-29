package range1

import "fmt"

func Demo1() {
	cities := []string{"Istanbul", "Ankara", "Antalya", "Bursa"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	/* As you can see this for loop won't work because we won't mention "i" because this isn't important for our duty.
	for i,city:= range cities{
		fmt.Println(cities)
	}
	*/

	//So that, we can use this for operation for our purpose.
	for _, city := range cities {
		fmt.Println(city)
	}
}
