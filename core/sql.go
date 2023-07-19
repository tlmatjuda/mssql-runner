package core

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tlmatjuda/this-and-that/files"
	"github.com/tlmatjuda/this-and-that/logs"
	"github.com/tlmatjuda/this-and-that/text"
	"log"
	"strings"
)

const MS_SQL_GO_KEYWORD = "GO"
const MS_SQL_FILE_EXTENSION = ".sql"

var dataAccess *sql.DB

func init() {

}

func RunSqlFile(sqlFilePath string) {

	if DatabaseConnectionIsAlive() &&
		text.StringNotBlank(sqlFilePath) &&
		files.Exists(sqlFilePath) &&
		strings.Contains(sqlFilePath, MS_SQL_FILE_EXTENSION) {

		sqlFileContent := files.ReadContent(sqlFilePath)
		sqlFileContent = strings.Replace(sqlFileContent, MS_SQL_GO_KEYWORD, "", -1)
		_, sqlError := dataAccess.Exec(sqlFileContent)
		if sqlError != nil {
			logs.Error.Fatalf("RunSqlFile() SQL ERROR : %v", sqlError.Error())
		}
	}

	defer func(dataAccess *sql.DB) {
		sqlError := dataAccess.Close()
		if sqlError != nil {
			logs.Error.Fatalf("RunSqlFile() SQL ERROR While closing connection : %v", sqlError.Error())
		}
	}(dataAccess)
}

func ConnectToDatabase(currenEnv DatabaseEnvironment) {

	var sqlError error
	// Create connection string
	connectionString := BuildConnectionString(currenEnv)
	dataAccess, sqlError = sql.Open("sqlserver", connectionString)
	if sqlError != nil {
		log.Fatalf("ConnectToDatabase() Error connecting to the database : %v", sqlError.Error())
	}

	log.Printf("Connected to the database successfully!")
}

func BuildConnectionString(currenEnv DatabaseEnvironment) string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		currenEnv.Host, currenEnv.Username, currenEnv.Password, currenEnv.Port)
}
func DatabaseConnectionIsAlive() bool {
	var canConnect = true
	context := context.Background()

	// Check if dataAccess is alive.
	err := dataAccess.PingContext(context)
	if err != nil {
		logs.Error.Printf("BuildConnectionString() Cannot connect to the database : %v", err.Error())
		canConnect = false
	}

	return canConnect
}
