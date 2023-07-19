package main

import (
	"fmt"
	"github.com/tlmatjuda/this-and-that/logs"
	"github.com/tlmatjuda/this-and-that/text"
	"os"
)

func main() {

	// Let's get all args from index 1 since the first arg is the script that we are executing
	var args = os.Args[1:]

	// We are expecting at least the first arg which is the main command
	if len(args) < 1 {
		logs.Error.Fatal("Not enough commands to run.")
	}

	environmentArg := args[0]
	sqlDirArg := args[1]
	validateArgs(environmentArg, sqlDirArg)

	var selectedEnvironment = findSelectedEnvironment(environmentArg)
	logs.Info.Printf("You are about to execute SQL files in no particular order towards the : [ %v ] environment", selectedEnvironment.Environment)
	logs.Info.Println("The database details are as follows : ")
	logs.Info.Printf("HOST : %v", selectedEnvironment.Host)
	logs.Info.Printf("PORT : %v", selectedEnvironment.Port)
	logs.Info.Printf("USER : %v", selectedEnvironment.Username)
	logs.Info.Printf("PASSWORD : ( Yea right :) )")
	logs.Info.Printf("")

	userConfirmationsArg := promptUser("If this is correct, please type either : Yes or No to continue ...")
	if text.EqualsIgnoreCase(KEY_YES, userConfirmationsArg) {
		// TODO : Run SQL here
	}

	logs.Info.Printf("All done with the process now.")
}

func promptUser(promptMessage string) string {
	var userInput string

	if text.StringNotBlank(promptMessage) {
		logs.Info.Println(promptMessage)
	}

	fmt.Scanln(&userInput)
	return userInput
}
