package functions

import "fmt"

// I learned different usage types of function for Golang:
func SaytoHi(username string) string {

	fmt.Println("Hi to my GOLang learning journey!", username)

	return username
}

func Addition(number1 int, number2 int) {

	fmt.Println(number1 + number2)

	//or

	equal := number1 + number2
	fmt.Println(equal)
}

//We use like that in real life.-->

func AdditionRealLife(number3 int, number4 int) int {

	equall := number3 + number4

	return equall
}
