package configs

import (
	"os"

	"github.com/gh0stl1m/notification-service/utils/constants"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type MongoDBConfig struct {
	URL      string
	Database string
}

func ReadMongoDBConfig() *MongoDBConfig {

	urlValue := getEnvAsString(constants.EnvironmentVariables.MONGO_URL)
	databaseValue := getEnvAsString(constants.EnvironmentVariables.MONGO_DB)

	return &MongoDBConfig{
		URL:      urlValue,
		Database: databaseValue,
	}

}

func getEnvAsString(name string) string {

	err := godotenv.Load()

	if err != nil {

		log.Fatalf("Error reading variable %s", name)
	}

	envVar := os.Getenv(name)

	if len(envVar) == 0 {
		log.Fatalf("Environment %s is not set", name)
	}

	return envVar
}
