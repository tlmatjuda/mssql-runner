package main

import (
	"fmt"
	"github.com/tlmatjuda/this-and-that/logs"
)

func main() {
	logs.Info.Println("Enter your argument here :  ")

	// var then variable name then variable type
	var userConfirmationArg string

	// Taking input from user
	fmt.Scanln(&userConfirmationArg)

	logs.Info.Printf("So your argument is : %v ", userConfirmationArg)

}
