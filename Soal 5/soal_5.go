package main

import (
	"fmt"
	"strings"
)


func main()  {
	var input string

	fmt.Print("Enter a string to check palindrom : ")
	fmt.Scanln(&input)
	input = strings.ToLower(input)

	sliceInput := strings.Split(input, "")
	sliceOutput := make([]string,len(input))
	for i := 0; i < len(input); i++ {
		sliceOutput[i] = sliceInput[len(input)-(i+1)]
	}

	output := strings.Join(sliceOutput, "")
	fmt.Println("Palindrom : ", input == output)
}