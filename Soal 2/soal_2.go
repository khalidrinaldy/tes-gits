package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main()  {
	fmt.Print("Enter email to check : ")
	var email string
	isValid := true

	fmt.Scanln(&email)
	
	//Contains @
	if !strings.Contains(email, "@") {
		isValid = false
	}

	//@.
	atSignDot, _ := regexp.MatchString(`(@)(\.)`,email)
	if !atSignDot {
		isValid = false
	}

	//Max 20 char before @
	if strings.Index(email, "@") > 20 {
		isValid = false
	}

	//Domain
	domain := strings.Split(email, "@")[1]
	if domain==".co.id" || domain==".id" {
		isValid = true
	} else {
		isValid = false
	}

	fmt.Println("Validitas email : ",isValid)
}