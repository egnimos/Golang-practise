/**********Addition of two number***************/

///What You Learn From This --->
///1) short declaration ":="
///2) how to create function
///3) what are parameters """functionName(parameter1 dataType, parameter2 dataType and soooo ooonnn)"""
///4) what are arguments
///5) assignment operator "="
///6) what are escape character "\n"
///7) what are verbs "%d, %s etc"
///8) return type "int, float, string"
///9) local scope and global scope variable


package practice

// import "fmt"


var globalVarable string

//add the two numbers
func addition(a int, b int) int {
	result := a + b
	return result
}

func subtraction(a int, b int) int {
	result := a - b
	return result
}

func division(a int, b int) int {
	result := a / b
	return result
}

func multiplication(a int, b int) int {
	result := a * b
	return result
}

// func intialize() (int, float64) {
// 	var a int = 20
// 	var b int = 10

// 	return a+b, 100.0
// }

// func main() {
// 	var localVariable string
// 	//addition
// 	result := addition(10, 10)
// 	fmt.Printf("your addition result is => %d\n", result)
// 	//subtraction
// 	result1 := subtraction(30, 40)
// 	fmt.Printf("your subtraction result1 is => %d\n", result1)
// 	//division
// 	result = division(200, 10)
// 	fmt.Printf("this is my division result %d\n", result)
// 	//multiplication
// 	result = multiplication(200, 10)
// 	fmt.Printf("this is my multiplication result %d", result)

// 	//initalize with some value
// 	globalVarable = "this is my global variable"
// 	localVariable = "this is my local variable"

// 	fmt.Printf("global variable result => %s", globalVarable)
// 	fmt.Printf("local variable result => %s", localVariable)
// }
