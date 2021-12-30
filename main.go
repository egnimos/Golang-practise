package main

import (
	"fmt"

	"github.com/Hanishkk/Golang-practise/practice_1"
)

func main() {
	var variable1 int
	var variable2 string
	var variable3 string
	var variable int
	//call the function
	//simple function call
	practice_1.FunctionName()
	//function with parameter call
	practice_1.FunctionWithParameter(12, 23.4, "hello")
	//function with parameter and return type call
	variable2, variable1, variable3 = practice_1.FunctionWithParameterAndReturnType(20)
	//function with only return type
	variable = practice_1.FunctionWithReturnType()

	fmt.Println(variable1, variable2, variable3, variable)
}