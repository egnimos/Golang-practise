package main

import "fmt"

var GlobalVariable string = "hello global"
var globalVaribale string = "hola amigous"

func main() {
	var localVariable int

	for i := 0; i < 10; i++ {
		localVariable = i
	}

	scope()

	fmt.Println(localVariable)
	fmt.Println(globalVaribale, GlobalVariable)
}

func scope() {
	fmt.Println(globalVaribale, GlobalVariable)
}