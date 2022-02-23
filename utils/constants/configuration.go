package constants

type environment struct {
	MONGO_URL string
	MONGO_DB  string
}

var EnvironmentVariables environment = environment{
	MONGO_URL: "MONGO_URL",
	MONGO_DB:  "MONGO_DB",
}
