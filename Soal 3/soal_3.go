package main

import (
	"fmt"
	"strings"
	"strconv"
)


func main()  {
	var inputClock string = "06:12:00 PM"
	hour, err := strconv.Atoi(string(inputClock[0])+string(inputClock[1]))
	minute, err := strconv.Atoi(string(inputClock[3])+string(inputClock[4]))
	_ = err

	if minute>60 {
		minute = minute-60
		hour = hour+1
	}

	if strings.Contains(inputClock, "PM") && hour<12{
		hour = hour+12
	} else if strings.Contains(inputClock, "AM") && hour>11 {
		hour = hour-12
	}

	var outputHour string = strconv.Itoa(hour)
	var outputMinute string= strconv.Itoa(minute)
	
	if hour<10 {
		outputHour = "0"+strconv.Itoa(hour)
	}
	if minute<10 {
		outputMinute = "0"+strconv.Itoa(minute)
	}

	var outputClock string = outputHour+":"+outputMinute
	fmt.Println("Input 12Hour Clock : ", inputClock)
	fmt.Print("24Hour Clock : ", outputClock)
}