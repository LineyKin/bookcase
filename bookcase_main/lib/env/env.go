package env

import (
	"os"

	env "github.com/joho/godotenv"
)

const port_key string = "PORT"
const sqlite_db_key string = "SQLITE_DBFILE"
const kafka_port_key = "KAFKA_PORT"

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

func GetPort() string {
	return getByKey(port_key)
}

func GetDbName() string {
	return getByKey(sqlite_db_key)
}

func GetKafkaPort() string {
	return getByKey(kafka_port_key)
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
