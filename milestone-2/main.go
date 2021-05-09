package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"milestone-2/config"
	"milestone-2/events"
	"milestone-2/models"
	"milestone-2/publisher"
	"github.com/google/uuid"
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
	var err error
	var order = models.Order{
		ID: uuid.New(),
	}

	var event = events.OrderReceived{
		EventBase: events.BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: order,
	}

	if err = publisher.PublishEvent(event, config.OrderReceivedTopicName); err != nil {
		log.WithField("orderID", order.ID).Error(err.Error())
	} else {
		log.WithField("event", event).Info("published event")
	}
}
