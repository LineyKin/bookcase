package app

import (
	"bookcase_log/internal/kafka/consumer"
	"bookcase_log/internal/storage"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Storage       *storage.Storage
	kafkaConsumer *consumer.KafkaConsumer
}

func New(db *sql.DB) (*App, error) {
	a := &App{}

	// слой хранилища
	a.Storage = storage.New(db)

	// брокер сообщений кафка (получатель)
	kc, err := consumer.New()
	if err != nil {
		log.Println("can't make kafka consumer: ", err)
		return nil, err
	}

	a.kafkaConsumer = kc

	return a, nil
}

func (a *App) Run() error {

	msgCnt := 0

	consumer, err := a.kafkaConsumer.Partition()
	if err != nil {
		log.Println("kafka partition error: ", err)
		return err
	}

	log.Println("Consumer successfully started")

	// 2. Handle OS signals - used to stop the process.
	sigchan := make(chan os.Signal, 1)

	// SIGINT - Сигнал прерывания (Ctrl-C) с терминала
	// SIGTERM - Сигнал завершения (сигнал по умолчанию для утилиты kill)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 3. Create a Goroutine to run the consumer / worker.
	doneCh := make(chan struct{})
	go func() {
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
	}()

	<-doneCh
	fmt.Println("Processed", msgCnt, "messages")

	// 4. Close the consumer on exit.
	if err := a.kafkaConsumer.Close(); err != nil {
		return err
	}

	return nil
}
