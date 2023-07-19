package main

import (
	"fmt"
	"github.com/tlmatjuda/mssql-runner/core"
	"github.com/tlmatjuda/this-and-that/files"
	"github.com/tlmatjuda/this-and-that/logs"
	"github.com/tlmatjuda/this-and-that/text"
	"os"
	"strings"
)

func main() {

	// Let's get all args from index 1 since the first arg is the script that we are executing
	var args = os.Args[1:]

	// We are expecting at least the first arg which is the main command
	if len(args) != core.CLI_ARGS_TOTAL {
		logs.Error.Fatal("There are arguments missing, please check")
	}

	environmentArg := args[core.CLI_ARGS_ENVIRONMENT_INDEX]
	sqlDirArg := args[core.CLI_ARGS_SQL_FILES_INDEX]
	core.ValidateArgs(environmentArg, sqlDirArg)

	var selectedEnvironment = core.FindSelectedEnvironment(environmentArg)
	logs.Info.Printf("You are about to execute SQL files in no particular order towards the : [ %v ] environment", selectedEnvironment.Environment)
	logs.Info.Printf("")
	logs.Info.Println("The database details are as follows : ")
	logs.Info.Printf("HOST : %v", selectedEnvironment.Host)
	logs.Info.Printf("PORT : %v", selectedEnvironment.Port)
	logs.Info.Printf("USER : %v", selectedEnvironment.Username)
	logs.Info.Printf("PASSWORD : ( Yea right :) )")
	logs.Info.Printf("")

	userConfirmationsArg := PromptUserInput("If this is correct, please type either : Yes or No to continue ...")
	core.ValidateConfirmationArg(userConfirmationsArg)
	if text.EqualsIgnoreCase(core.KEY_YES, userConfirmationsArg) {

		// Connect to the database first
		core.ConnectToDatabase(selectedEnvironment)

		// Run SQL Files one by one
		filesList := files.List(sqlDirArg)
		for _, fileData := range filesList {
			if strings.Contains(fileData.Name(), core.MS_SQL_FILE_EXTENSION) {
				sqlFilePath := sqlDirArg + "/" + fileData.Name()
				logs.Info.Printf("Running SQL File : %v", sqlFilePath)
				core.RunSqlFile(sqlFilePath)
			}
		}
	} else {
		logs.Info.Printf("")
		logs.Info.Println("Since you did not type : \"Yes\" we will not continue with the process, bye bye! ")
		os.Exit(-1)
	}

	logs.Info.Printf("Process complete")
}

func PromptUserInput(promptMessage string) string {
	var userInput string

	if text.StringNotBlank(promptMessage) {
		logs.Info.Println(promptMessage)
	}

	fmt.Scanln(&userInput)
	return userInput
}
