package main

import "fmt"

func main() {
	var age = 20
	
	var eext = "John"
	var text = "john"
	fmt.Println("Before:", eext, age)

	update(&age, &eext)

	fmt.Println("After :", eext, age)
	fmt.Println(text)
}
func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	eext= eext + "haha"
	
}