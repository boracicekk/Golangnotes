package loops

import "fmt"

func Demo1() {
	var sayHi string = "Hello World!"
	count := 5

	for i := 0; i <= count; i++ {
		fmt.Println(sayHi)
	}

	//infinite loop
	/*
		i := 0

		for i <= count {
			fmt.Println(sayHi)
		}

	*/

	//other for loop syntax:

	/*	for i<=count {
			fmt.Println(sayHi)
			i=i+1
		}
	*/
	fmt.Println("Done.")
}
