package app

import (
	"bookcase/internal/handlers"
	"bookcase/internal/kafka"
	"bookcase/internal/service"
	"bookcase/internal/storage"
	"bookcase/lib/env"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	kp      *kafka.Producer
	storage *storage.Storage
	serv    *service.Service
	hand    *handlers.Handlers

	gin *gin.Engine
}

func New(db *sql.DB, kp *kafka.Producer) (*App, error) {
	a := &App{}

	// кафка
	a.kp = kp

	// слой хранилища
	a.storage = storage.New(db)

	// слой сервиса
	a.serv = service.New(a.storage)

	// слой эндпоинтов
	a.hand = handlers.New(a.serv, a.kp)

	// роутер
	a.gin = gin.Default()

	// ручка для главной страницы
	a.gin.GET("/", a.hand.FileServer)
	a.gin.Static("/style", "./web/style")
	a.gin.Static("/js", "./web/js")

	// ручка добавления авторов
	a.gin.POST("api/author/add", a.hand.AddAuthor)

	// ручка для выгрузки списка книг
	a.gin.GET("api/book/list", a.hand.GetBookList)

	// ручка для выгрузки количества книг
	a.gin.GET("api/book/count", a.hand.GetBookCount)

	// ручка добавления книги
	a.gin.POST("api/book/add", a.hand.AddBook)

	// ручка выгрузки авторов для подсказки в форме добавления книги
	a.gin.GET("api/author/hint", a.hand.GetAuthorList)

	// ручка выгрузки списка издательств
	a.gin.GET("api/publishingHouse/list", a.hand.GetPublishingHouseList)

	return a, nil
}

func (a *App) Run() error {
	port := env.GetPort()
	fmt.Printf("http://localhost:%s/\n", port)
	err := a.gin.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	return nil
}
