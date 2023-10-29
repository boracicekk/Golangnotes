// Variadic function basic:
package functions

func AdditionVariadicFunction(number ...int) int {
	additionvariadic := 0
	for i := 0; i < len(number); i++ {
		additionvariadic = additionvariadic + number[i]
	}
	return additionvariadic
}
