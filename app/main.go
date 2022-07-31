package main

import (
	"bitbucket.org/ripleyx/import-service/app/import/application/split"
	"bitbucket.org/ripleyx/import-service/app/import/infrastructure/queue/kafka/common"
	"bitbucket.org/ripleyx/import-service/app/import/infrastructure/repository/cloudstorage"
	"bitbucket.org/ripleyx/import-service/app/import/infrastructure/subscriber"
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/bus/inmemory"
	"bitbucket.org/ripleyx/import-service/app/shared/infrastructure/rest"
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	serverHost = os.Getenv("SERVER_HOST")
	serverPort = os.Getenv("SERVER_PORT")

	bucket = os.Getenv("IMPORT_BUCKET")

	kafkaUsername = os.Getenv("KAFKA_USERNAME")
	kafkaPassword = os.Getenv("KAFKA_PASSWORD")
	kafkaBrokers  = strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	kafkaGroupId  = os.Getenv("KAFKA_GROUP_ID")

	newImportTopic = os.Getenv("KAFKA_NEW_IMPORT_TOPIC")
)

func main() {
	server := rest.New()
	objectRepository := cloudstorage.NewGetObjectRepository(bucket)
	splitService := split.NewSplitService(objectRepository)
	splitCommandHandler := split.NewSplitCommandHandler(splitService)
	commandBus := inmemory.NewCommandBusMemory()
	commandBus.Register(split.ImportSplitType, splitCommandHandler)

	splitSubscriber := subscriber.NewSubscriberSplitter(
		kafkaGroupId,
		newImportTopic,
		common.GetDialer(kafkaUsername, kafkaPassword),
		kafkaBrokers...,
	)
	ctxSplitSubscriber, cancelSplitSubscriber := context.WithCancel(context.Background())
	defer cancelSplitSubscriber()
	go splitSubscriber.ReadMessage(ctxSplitSubscriber)

	go func() {
		if err := server.StartServer(rest.Setup(serverHost, serverPort)); err != http.ErrServerClosed {
			server.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctxServer, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxServer); err != nil {
		server.Logger.Fatal(err)
	}
}
