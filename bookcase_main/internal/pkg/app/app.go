package app

import (
	"bookcase/internal/db"
	"bookcase/internal/handlers"
	"bookcase/internal/kafka"
	"bookcase/internal/service"
	"bookcase/internal/storage"
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

func New(appDB db.AppDB, kp *kafka.Producer) (*App, error) {
	a := &App{}

	// кафка
	a.kp = kp

	// слой хранилища
	a.storage = storage.New(appDB)

	// слой сервиса
	a.serv = service.New(a.storage)

	// слой эндпоинтов
	a.hand = handlers.New(a.serv, a.kp)

	// роутер
	a.gin = gin.Default()

	// ручка для главной страницы
	a.gin.GET("/", a.hand.AuthMW(), a.hand.FileServer)

	// ручка страницы регистрации
	a.gin.GET("/auth", a.hand.FileServer)

	// ручка страницы общего списка книг
	a.gin.GET("/total", a.hand.AuthMW(), a.hand.FileServer)

	a.gin.Static("/style", "./web/style")
	a.gin.Static("/js", "./web/js")

	// ручка регистрации пользователя
	a.gin.POST("register", a.hand.Register, a.hand.LogMW())

	// ручка входа пользователя
	a.gin.POST("login", a.hand.Login, a.hand.LogMW())

	// ручка добавления авторов
	a.gin.POST(
		handlers.ADD_AUTHOR_URL,
		a.hand.AuthMW(),
		a.hand.AddAuthor,
		a.hand.LogMW(),
	)

	// ручка для выгрузки списка книг пользователя
	a.gin.GET("api/book/list", a.hand.AuthMW(), a.hand.GetBookList)

	// ручка для выгрузки списка книг пользователя
	a.gin.GET("api/book/list/total", a.hand.AuthMW(), a.hand.GetBookListTotal)

	// ручка для выгрузки количества книг пользователя
	a.gin.GET("api/book/count", a.hand.AuthMW(), a.hand.GetBookCount)

	// ручка для выгрузки количества книг всех пользователей
	a.gin.GET("api/book/count/total", a.hand.GetBookCountTotal)

	// ручка добавления книги
	a.gin.POST(
		handlers.ADD_BOOK_URL,
		a.hand.AuthMW(),
		a.hand.AddBook,
		a.hand.LogMW(),
	)

	// ручка выгрузки авторов для подсказки в форме добавления книги
	a.gin.GET("api/author/hint", a.hand.AuthMW(), a.hand.GetAuthorList)

	// ручка выгрузки списка издательств
	a.gin.GET("api/publishingHouse/list", a.hand.AuthMW(), a.hand.GetPublishingHouseList)

	return a, nil
}

func (a *App) Run() error {

	err := a.gin.Run(":1991")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	return nil
}
