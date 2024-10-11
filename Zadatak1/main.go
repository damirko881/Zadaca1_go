/*1. Declare two integers, `firstNumber` and `secondNumber`, assign values 20 and 30 to them.
Swap values of `firstNumber` and `secondNumber` without using third variable
After all, print values of `firstNumber` and `secondNumber`.*/

package main

import "fmt"

func main() {
	var firstNumber int = 20
	var secondNumber int = 30

	fmt.Println("Prvi broj: ", firstNumber)
	fmt.Println("Drugi broj: ", secondNumber)

	firstNumber = firstNumber + secondNumber
	secondNumber = firstNumber - secondNumber
	firstNumber = firstNumber - secondNumber

	fmt.Println("Prvi broj nakon zamjene: ", firstNumber)
	fmt.Println("Drugi broj nakon zamjene: ", secondNumber)

}
