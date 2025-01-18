package app

import (
	"bookcase_log/internal/handlers"
	"bookcase_log/internal/kafka/consumer"
	"bookcase_log/internal/storage"
	"bookcase_log/lib/env"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

type App struct {
	Storage       *storage.Storage
	kafkaConsumer *consumer.KafkaConsumer
	hand          *handlers.Handlers

	gin *gin.Engine
}

func New(db *sql.DB) (*App, error) {
	a := &App{}

	// слой хранилища
	a.Storage = storage.New(db)

	// слой эндпоинтов
	a.hand = handlers.New()

	// роутер
	a.gin = gin.Default()

	// ручка для главной страницы
	a.gin.GET("/", a.hand.FileServer)

	// брокер сообщений кафка (получатель)
	kc, err := consumer.New()
	if err != nil {
		log.Println("can't make kafka consumer: ", err)
	}

	a.kafkaConsumer = kc

	return a, nil
}

func (a *App) Run() error {

	if a.kafkaConsumer != nil {
		go runKafkaConsumer(a.kafkaConsumer)
	}

	port := env.GetPort()
	fmt.Printf("http://localhost:%s/\n", port)
	err := a.gin.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	return nil
}

func runKafkaConsumer(kc *consumer.KafkaConsumer) {
	msgCnt := 0

	consumer, err := kc.Partition()
	if err != nil {
		log.Println("kafka partition error: ", err)
	}

	log.Println("Consumer successfully started")

	// 2. Handle OS signals - used to stop the process.
	sigchan := make(chan os.Signal, 1)

	// SIGINT - Сигнал прерывания (Ctrl-C) с терминала
	// SIGTERM - Сигнал завершения (сигнал по умолчанию для утилиты kill)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 3. Create a Goroutine to run the consumer / worker.
	doneCh := make(chan struct{})
	//go func() {
	for {
		select {
		case err := <-consumer.Errors():
			fmt.Println(err)
		case msg := <-consumer.Messages():
			msgCnt++
			fmt.Printf("Received order Count %d: | Topic(%s) | Message(%s) \n", msgCnt, string(msg.Topic), string(msg.Value))
			order := string(msg.Value)
			fmt.Printf("Добавлен новый автор: %s\n", order)
		case <-sigchan:
			fmt.Println("Interrupt is detected")
			doneCh <- struct{}{}
		}
	}
	//}()

	<-doneCh
	fmt.Println("Processed", msgCnt, "messages")

	// 4. Close the consumer on exit.
	if err := kc.Close(); err != nil {
		log.Println("can't close consumer: ", err)
	}

}
