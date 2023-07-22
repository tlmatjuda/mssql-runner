package core

import (
	"fmt"
	"github.com/tlmatjuda/this-and-that/files"
	"github.com/tlmatjuda/this-and-that/logs"
	"github.com/tlmatjuda/this-and-that/text"
	"os"
)

func PromptUserInput(promptMessage string) string {
	var userInput string

	if text.StringNotBlank(promptMessage) {
		logs.Info.Println(promptMessage)
	}

	fmt.Scanln(&userInput)
	return userInput
}

func RunSQLFiles(sqlFiles []string, selectedEnvironment DatabaseEnvironment) {
	if len(sqlFiles) > 0 {
		logs.Info.Printf("Found %v SQL files", len(sqlFiles))
		// Connect to the database first
		ConnectToDatabase(selectedEnvironment)

		// Run the SQL files.
		for _, sqlFile := range sqlFiles {
			logs.Info.Printf("Running SQL File : %v", sqlFile)
			RunSqlFile(sqlFile)
		}

		// Now Closed Connection since we are done.
		CloseDatabaseConnection()
	} else {
		logs.Info.Printf("No \"%v\" files were found", MS_SQL_FILE_EXTENSION)
	}
}

func FindSqlFilesInDirectory(sqlDirArg string) []string {
	filterSuffix := DELIMTER_ASTERIX + MS_SQL_FILE_EXTENSION
	sqlFiles := func() []string {
		if DELIMTER_PERIOD == sqlDirArg {
			presentWorkingDirectory, _ := os.Getwd()
			return files.ListByWildcard(presentWorkingDirectory, filterSuffix)
		} else {
			return files.ListByWildcard(sqlDirArg, filterSuffix)
		}
	}()
	return sqlFiles
}
