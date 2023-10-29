package maps

import "fmt"

func Demo1() {
	//key-value
	dictionary := make(map[string]string)
	dictionary["table"] = "Masa"
	dictionary["book"] = "Kitap"
	dictionary["pencil"] = "Kalem"

	//How we can print our key's value?
	fmt.Println(dictionary["book"])

	//The map structure acts like an array in the background.
	fmt.Println("Number of dictionary element:", len(dictionary))

	//we can delete with delete function in our map structure.
	delete(dictionary, "book")
	//You can check:
	fmt.Println("Number of dictionary element:", len(dictionary))
	fmt.Println("Dictionary format:", dictionary)

	//How we can define a value in our key element?

	value := dictionary["table"]
	fmt.Println(value)

	//Also:

	value2, IsThere := dictionary["pencil"]
	fmt.Println(value2)
	fmt.Println("Have we got:", IsThere)

	value3, IsThere2 := dictionary["lamb"]
	fmt.Println(value3)
	fmt.Println("Have we got:", IsThere2)

}
