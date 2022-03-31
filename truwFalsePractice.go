package main

import "fmt"

func compareValueHandler(userInput float64) string {
	var secretValue float64 = 69
	resultMsg := "Nice!"

	if userInput < secretValue {
		fmt.Println("higher")
	}
	if userInput > secretValue {
		fmt.Println("lower")
	}
	if userInput == secretValue {
		fmt.Println("Well done")
	}
	return resultMsg
}

func main() {
	var userInput float64
	fmt.Println("Please enter a value from 1 - 100")
	fmt.Scanln(&userInput)
	compareValueHandler(userInput)
}
