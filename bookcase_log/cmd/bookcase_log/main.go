package main

import (
	"bookcase_log/internal/db"
	"bookcase_log/internal/pkg/app"
	"log"
)

const APP = "bookcase_log"

func main() {

	log.Printf("%s start", APP)
	db, err := db.InitPostgresDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Printf("%s start database", APP)

	a, err := app.New(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s build", APP)

	err = a.Run()
	if err != nil {
		log.Println("runtime error: ", err)
		log.Fatal(err)
	}

	log.Printf("%s run", APP)

}
