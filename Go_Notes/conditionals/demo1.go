package conditionals

import "fmt"

// The basics "if condition" with example:
func Demo1() {
	var bankaccount1 float64 = 1000
	var withdrawMoney float64 = 1900

	if withdrawMoney > bankaccount1 {
		fmt.Println("ERROR! You haven't got enough money in your bank account..")
	} else if withdrawMoney == bankaccount1 {
		fmt.Println("Preparing your money.. ")
	} else {
		fmt.Println("Done.")
	}
}
