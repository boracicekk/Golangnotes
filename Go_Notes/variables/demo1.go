package variables

import "fmt"

func Demo1() {
	// Hi everyone!

	/*
		I will share my GOlang learning journey with you!

		I will try to show many type of usage about Golang basics.

		Lets start!

	*/
	//Print functions' usage:
	fmt.Println("Hello")
	fmt.Print("World!")

	//These page's main topic is declare variables. Lets see!

	var intagerNumbers int = 5
	fmt.Println(intagerNumbers)

	// Float has 2 variety these are float32 and float64. Its depends on your decimal number's bit number.
	var decimalNumbers float32 = 5.3
	fmt.Println(decimalNumbers)

	var decimaLNumbers float64 = 55555555555555555555555555555555555555555555555.33333333333333333333333333333333333333333333333333333333
	fmt.Println(decimaLNumbers)

	// You can check to switch the float32 and float64's values. You will see the error message.

	var stringvariable string = "BoraÇiçek"
	fmt.Println(stringvariable)

	// Let's do some calculations using our integer values!

	var taxes int = 7
	salary := 34000 // GOLANG PROVİDE THİS UNIQUE DECLARE TYPE TO US! OK, HOW WE CAN SEE THE TYPE OF SALARY?
	fmt.Printf("Variable type : %T ", salary)
	var bankaccount int = salary - (salary * int(taxes) / 100)

	fmt.Println(bankaccount)

	//Also we can do some calculations in Print function without variables.
	fmt.Println(34000 - (34000 * 7 / 100)) // As you can see we can reach same result on our screen.

	//Logical data type
	var logic bool = false
	fmt.Println(logic)

	//Also, we can define according to condition!

	var logic1 string = "Bora"
	var logic2 string = "Çiçek"

	var condition bool = logic1 == logic2
	fmt.Println(condition)

	// or

	var condition2 bool = logic1 != logic2
	fmt.Println(condition2)

}
