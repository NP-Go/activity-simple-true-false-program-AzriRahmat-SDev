package main

import "fmt"

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessage := "Done"
	secretValue := 88

	//Insert your code from here

	if value < secretValue {
		fmt.Println("Too low value")
	}
	if value > secretValue {
		fmt.Println("Too high value")
	}
	if value == secretValue {
		fmt.Println("well done")
	}
	//do not remove this line
	return resultMessage
}

func sample() {
	var guess int
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&guess)
	compare(guess)
}
