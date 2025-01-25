package main

import (
	"bookcase/internal/db"
	"bookcase/internal/kafka"
	"bookcase/internal/pkg/app"
	"log"
)

func main() {

	// 1. подключение к БД
	appDB, err := db.Init(db.PG_DRIVER) // можно переключиться на sqlite, аргумент db.SQLITE_DRIVER
	if err != nil {
		log.Fatal("can't connect to db: ", err)
	}
	defer appDB.Connection.Close()

	// 2. кафка (продюсер) для слоя эндпоинтов
	kp := kafka.NewProducer()
	defer kp.Close()

	// 3. объект всего приложения
	a, err := app.New(appDB, kp)
	if err != nil {
		log.Fatal("can't build app: ", err)
	}

	// 4. запуск приложения
	err = a.Run()
	if err != nil {
		log.Fatal("can't run app: ", err)
	}
}
