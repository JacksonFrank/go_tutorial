package main

import "fmt"

func main() {
	var name string = "Jackson"

	var message string
	message = "hello"

	var number uint32 = 260

	var number_implicit = 350

	other_number_implicit := 11.57

	fmt.Println(message, name, "number:", number)

	fmt.Printf("implicit number type: %T\nohter implcit number type: %T\n", number_implicit, other_number_implicit)

	var default_number uint32
	fmt.Println("default uin32:", default_number)
}
