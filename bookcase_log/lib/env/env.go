package env

import (
	"os"

	env "github.com/joho/godotenv"
)

const portKey string = "PORT"
const kafkaPortPey = "KAFKA_PORT"
const pgHostKey = "PG_HOST"
const pgPortKey = "PG_PORT"
const pgUserKey = "PG_USER"
const pgPasswordKey = "PG_PASSWORD"
const pgDbNameKey = "PG_DBNAME"

func getByKey(key string) string {
	err := env.Load(".env")

	// костыль для тестов
	if err != nil {
		err = env.Load("../.env")
	}

	if err != nil {
		panic("Невозможно загрузить .env")
	}

	return os.Getenv(key)
}

func GetPgDbName() string {
	return getByKey(pgDbNameKey)
}

func GetPgPassword() string {
	return getByKey(pgPasswordKey)
}

func GetPgUser() string {
	return getByKey(pgUserKey)
}

func GetPgPort() string {
	return getByKey(pgPortKey)
}

func GetPgHost() string {
	return getByKey(pgHostKey)
}

func GetPort() string {
	return getByKey(portKey)
}

func GetKafkaPort() string {
	return getByKey(kafkaPortPey)
}
