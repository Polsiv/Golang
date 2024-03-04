package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println("Hello World!")

	//integers
	var intNum int = 17
	fmt.Println(intNum)

	//floats
	var floatNum float64 = 18.2
	fmt.Println(floatNum)

	var floatResult float64 = float64(intNum) + floatNum
	fmt.Println(floatResult)
	var intNum1 int = 1
	var intNum2 int = 2

	//basic operators
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	//strings
	var myString string = "Hello bitch"
	var myString2 string = "hello" + " " + "bIthc"
	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(utf8.RuneCountInString(myString))
	fmt.Println(len(myString)) //returns the number of bytes

	//booleans
	var myBool bool = true
	fmt.Println(myBool)

	//shortcuts
	var myNumber = 1
	fmt.Println(myNumber)
	myNumber2 := 2
	fmt.Println(myNumber2)

	//multi assigment
	myNumber3, myNumber4 := 3, 4
	fmt.Println(myNumber3 + myNumber4)

	//const
	const myContst int = 23
	fmt.Println(myContst)

	//functions
	printMe(myString)
	fmt.Printf("The result is %v in division\n", intDivision(intNum2, intNum1))

	//switch case
	var myvar1 int = 2

	switch myvar1 {
	case 0:
		fmt.Println("El pepe")
	case 1, 2:
		fmt.Println("el pepe2")
	}

	//arrrays
	var intArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)

	//loops

}

func printMe(printvalue string) {
	fmt.Println(printvalue)
}

// the (int, int) is the amount of returns
func intDivision(num int, den int) int {
	return num / den
}
