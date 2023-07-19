package core

import (
	"encoding/json"
	"github.com/tlmatjuda/this-and-that/files"
	"log"
	"os"
	"strconv"
)

const (
	MS_SQL_SERVER_PORT     = "MS_SQL_SERVER_PORT"
	MS_SQL_SERVER_HOST     = "MS_SQL_SERVER_HOST"
	MS_SQL_SERVER_USERNAME = "MS_SQL_SERVER_USERNAME"
	MS_SQL_SERVER_PASSWORD = "MS_SQL_SERVER_PASSWORD"
)

var (
	databaseEnvList  []DatabaseEnvironment
	databaseEnvNames []string
)

func init() {
	databaseEnvList = fetchDatabaseEnvironments()

	for _, envRecord := range databaseEnvList {
		databaseEnvNames = append(databaseEnvNames, envRecord.Environment)
	}
}

func fetchDatabaseEnvironments() []DatabaseEnvironment {
	content := loadJsonConfigContent()
	return jsonToEnvironmentStructure(content)
}

func loadJsonConfigContent() string {
	workingDirectory, _ := os.Getwd()
	return files.ReadContent(workingDirectory + "/configs/mssql-env-conf.json")
}

func jsonToEnvironmentStructure(jsonContent string) []DatabaseEnvironment {
	var environments []DatabaseEnvironment

	// Converting the string to bytes, for the JSON Marshaller.
	contentAsBytes := []byte(jsonContent)

	// Now ask the json util to marshal or deserialize for us into the LibraryRecord type.
	err := json.Unmarshal(contentAsBytes, &environments)
	if err != nil {
		log.Fatal(err)
	}

	return environments
}

func FindSelectedEnvironment(environmentArg string) DatabaseEnvironment {
	var currenEnv DatabaseEnvironment
	for _, envRecord := range databaseEnvList {
		if envRecord.Environment == environmentArg {
			currenEnv = envRecord
			break
		}
	}
	return currenEnv
}

func environmentExists(environmentName string) bool {
	var exists bool
	for _, envName := range databaseEnvNames {
		if environmentName == envName {
			exists = true
			break
		}
	}

	return exists
}

func environmentNotExists(environmentName string) bool {
	return !environmentExists(environmentName)
}

func buildLocalDatabaseEnvironment() DatabaseEnvironment {

	sqlPort, _ := strconv.Atoi(os.Getenv(MS_SQL_SERVER_PORT))

	return DatabaseEnvironment{
		Environment: "local",
		Username:    os.Getenv(MS_SQL_SERVER_USERNAME),
		Password:    os.Getenv(MS_SQL_SERVER_PASSWORD),
		Port:        sqlPort,
		Host:        os.Getenv(MS_SQL_SERVER_HOST),
	}
}

type DatabaseEnvironment struct {
	Environment string `json:"environment"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Port        int    `json:"port"`
	Host        string `json:"host"`
}
