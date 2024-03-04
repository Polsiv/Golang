package main

import (
	"bufio" //create buffer reader
	"fmt"
	"os" //open fule
	"strconv"
)

func removeIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func scanData(filename string) []int {
	file, ferr := os.Open(filename)
	var numbers []int

	scanner := bufio.NewScanner(file)
	if ferr != nil {
		panic(ferr)
	}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("error! omg")
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func checkNumbers(number int) string {

	var toComplete string

	if number%3 == 0 {
		toComplete += "Fizz"
	}
	if number%5 == 0 {
		toComplete += "Buzz"
	}
	if number%3 != 0 && number%5 != 0 {
		toComplete += strconv.Itoa(number)
	}
	return toComplete
}

func printResult(numList []int) {
	for i := range numList {
		fmt.Printf("%v %v\n", numList[i], checkNumbers(numList[i]))
	}
}

func main() {
	var numList []int = scanData("input.txt")
	var modNumList []int = removeIndex(numList, 0)
	printResult(modNumList)
}
