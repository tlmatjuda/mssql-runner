package core

import (
	"github.com/tlmatjuda/this-and-that/logs"
	"github.com/tlmatjuda/this-and-that/text"
)

func ValidateArgs(environmentArg string, sqlDirArg string) {
	if text.StringBlank(environmentArg) {
		logs.Error.Fatalln("Environment name required")
	}

	if text.StringBlank(sqlDirArg) {
		logs.Error.Fatalln("SQL file directory argument expected")
	}

	if environmentNotExists(environmentArg) {
		logs.Error.Fatalf("Environment name : %v does not exists, we only support : %v", environmentArg, databaseEnvNames)
	}

	if environmentNotExists(sqlDirArg) {
		logs.Error.Fatalf("SQL file directory : %v does not exists", sqlDirArg)
	}
}

func validateConfirmationArg(userConfirmationArg string) {
	if text.StringBlank(userConfirmationArg) {
		logs.Error.Fatalln("Confirmation response expected")
	}

	if text.NotEqualsIgnoreCase(KEY_YES, userConfirmationArg) && text.NotEqualsIgnoreCase(KEY_NO, userConfirmationArg) {
		logs.Error.Fatalln("Please type either : \"Yes\" or \"No\"")
	}
}
