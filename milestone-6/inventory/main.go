package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bpmericle/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/milestone6/code/config"
	"github.com/bpmericle/Asynchronous-Event-Handling-Using-Microservices-and-Kafka/milestone6/code/inventory/cmd/consumer"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(config.LogLevel())
}

func main() {
	startTime := time.Now()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.WithField("uptime", time.Since(startTime).String()).
			WithField("signal", sig.String()).
			Error("interrupt signal detected")
		os.Exit(0)
	}()

	c := consumer.Consumer{
		Broker: config.BrokerAddress(),
		Group:  config.ConsumerGroup(),
		Topic:  config.OrderReceivedTopicName,
	}

	log.Fatal(c.SubscribeAndListen())
}
